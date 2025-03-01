package mw

import (
	"context"
	"errors"
	"fmt"
	"github.com/PTS0118/go-mall/api/biz/dal/redis"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PTS0118/go-mall/api/biz/model"
	utils2 "github.com/PTS0118/go-mall/api/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddle   *jwt.HertzJWTMiddleware
	IdentityKey = "uid"
)

func JWTInit() {
	var err error
	JwtMiddle, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "hertz jwt",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			// 设置token到redis
			// 从 Payload 中解析用户 ID
			parsedToken, err := JwtMiddle.ParseTokenString(token)
			if err != nil {
				klog.Info("Token 解析失败: %v", err)
			} else {
				claims := jwt.ExtractClaimsFromToken(parsedToken)
				userID := claims[IdentityKey].(string)
				tokenKey := "token_" + userID
				// 存储到 Redis
				if err := redis.RedisClient.SetEx(ctx, tokenKey, token, 3600*time.Second).Err(); err != nil {
					klog.Info("Redis 存储失败: %v", err)
				}
			}

			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "success",
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format username'"`
				Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format password'"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := model.CheckUser(loginStruct.Username, utils2.MD5(loginStruct.Password))
			if err != nil {
				return nil, err
			}
			if len(users) == 0 {
				return nil, errors.New("user already exists or wrong password")
			}

			return &model.User{
				Base:     model.Base{Id: users[0].Base.Id},
				Username: users[0].Username,
				Email:    users[0].Email,
				Password: users[0].Password,
				Role:     users[0].Role,
			}, nil
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			// 验证token是否还在redis (不在说明已经登出,token不一致表示已经重新登录了)
			token := jwt.GetToken(ctx, c)
			claims := jwt.ExtractClaims(ctx, c)
			idFloat := claims[IdentityKey].(string)
			tokenKey := "token_" + idFloat
			if ok, _ := ValidateToken(tokenKey, token, ctx); !ok {
				c.AbortWithStatusJSON(401, utils.H{"error": "Token已失效，请重新登录"})
			}
			userId, _ := strconv.Atoi(idFloat)
			return &model.User{
				Base: model.Base{Id: userId},
				Role: claims["role"].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				//fmt.Printf("role:%+v", v)
				return jwt.MapClaims{
					IdentityKey: strconv.Itoa(v.Base.Id),
					"role":      v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			//获取到token
			authHeader := c.GetHeader("Authorization")
			token := strings.TrimPrefix(string(authHeader), "Bearer ")
			parsedToken, err := JwtMiddle.ParseTokenString(token)
			if err != nil {
				klog.Info("Token 解析失败: {}", err)
			} else {
				claims := jwt.ExtractClaimsFromToken(parsedToken)
				userID := claims[IdentityKey].(string)
				tokenKey := "token_" + userID
				// 将token加入黑名单
				if err := redis.RedisClient.SetEx(ctx, tokenKey, "", 1*time.Second).Err(); err != nil {
					klog.Info("Redis 设置过期失败: %v", err)
				}
			}
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": http.StatusOK,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}

func ValidateToken(tokenKey string, targetToken string, ctx context.Context) (bool, error) {
	// 1. 获取 Redis 值
	val, err := redis.RedisClient.Get(ctx, tokenKey).Result()

	// 2. 处理键不存在的情况
	if err != nil {
		return false, fmt.Errorf("Redis 查询失败: %v", err) // 其他错误（如连接问题）
	}

	// 3. 判断值是否为空字符串
	if val == "" {
		return false, nil
	}

	// 4. 验证值是否与目标 Token 一致
	return val == targetToken, nil
}
