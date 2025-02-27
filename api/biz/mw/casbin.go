package mw

import (
	"context"
	"fmt"
	"github.com/PTS0118/go-mall/api/biz/dal/mysql"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/jwt"
	"strings"
)

var Enforcer *casbin.Enforcer

// 初始化 Casbin Enforcer（带错误处理）
func InitCasbin() (*casbin.Enforcer, error) {
	// 2. 创建 Casbin 适配器
	// 使用 GORM 适配器（自动创建 casbin_rule 表）
	adapter, err := gormadapter.NewAdapterByDB(mysql.DB)
	if err != nil {
		return nil, err
	}

	// 3. 加载 RBAC 模型配置
	modelConfig := `
    [request_definition]
    r = sub, obj, act

    [policy_definition]
    p = sub, obj, act

    [policy_effect]
    e = some(where (p.eft == allow))

    [matchers]
    m = r.sub == p.sub && DoublestarMatch(p.obj, r.obj)  && (r.act == p.act || p.act == "*")
    `

	// 3. 加载模型和适配器
	m, err := model.NewModelFromString(modelConfig)
	if err != nil {
		return nil, err
	}

	// 4. 创建 Enforcer 实例
	Enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}

	// 注册自定义函数
	Enforcer.AddFunction("DoublestarMatch", DoublestarMatch)

	// 5. 配置增强选项
	Enforcer.EnableAutoSave(true) // 自动保存策略到数据库
	Enforcer.EnableLog(false)     // 启用日志（生产环境可关闭）

	// 加载现有策略
	if err = Enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	// 6. 初始化默认策略（可选）
	if hasPolicy, _ := Enforcer.HasPolicy("admin", "/**", "*"); !hasPolicy {
		if _, err = Enforcer.AddPolicy("admin", "/**", "*"); err != nil {
			return nil, err
		}
	}

	if hasPolicy, _ := Enforcer.HasPolicy("user", "/a/**", "*"); !hasPolicy {
		if _, err = Enforcer.AddPolicy("user", "/a/**", "*"); err != nil {
			return nil, err
		}
	}

	return Enforcer, nil
}

func CasbinMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 提取 claims（带错误处理）
		claims, exists := c.Get("JWT_PAYLOAD")
		if !exists {
			c.AbortWithStatusJSON(401, "未提供有效 Token")
			return
		}

		// 类型断言保护
		jwtClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(401, "Token 格式错误")
			return
		}

		// 检查 role 字段是否存在且类型正确
		role, ok := jwtClaims["role"].(string)
		if !ok || role == "" {
			c.AbortWithStatusJSON(403, "用户角色未定义")
			return
		}

		// 执行权限检查
		// 去除查询参数
		path := strings.Split(string(c.Request.URI().Path()), "?")[0]
		method := string(c.Request.Method())
		klog.Info("role:", role, " path:", path, " method:", method)

		// 打印当前所有策略
		//policies, _ := Enforcer.GetPolicy()
		//fmt.Printf("[Casbin] 当前策略列表: %v", policies)

		ok, _ = Enforcer.Enforce(role, path, method)
		if !ok {
			c.AbortWithStatusJSON(403, "权限不足")
			return
		}
		c.Next(ctx)
	}
}

func DoublestarMatch(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return false, fmt.Errorf("DoublestarMatch requires exactly two arguments")
	}

	pattern, ok1 := args[0].(string)
	path, ok2 := args[1].(string)

	if !ok1 || !ok2 {
		return false, fmt.Errorf("DoublestarMatch requires string arguments")
	}

	matched, err := doublestar.Match(pattern, path)
	if err != nil {
		return false, err
	}
	return matched, nil
}
