package mw

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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
			claims := jwt.ExtractClaims(ctx, c)
			idFloat := claims[IdentityKey].(float64)
			return &model.User{
				Base: model.Base{Id: int(idFloat)},
				Role: claims["role"].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				fmt.Printf("role:%+v", v)
				return jwt.MapClaims{
					IdentityKey: v.Base.Id,
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
	})
	if err != nil {
		panic(err)
	}
}
