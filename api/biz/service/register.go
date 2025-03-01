package service

import (
	"context"

	auth "github.com/PTS0118/go-mall/api/hertz_gen/api/auth"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpcuser "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

// @Summary 用户注册
// @Description 通过RPC调用用户注册
// @Tags Auth
// @Accept json
// @Produce json
// @Param req body auth.RegisterReq true "用户注册请求"
// @Success 200 {object} auth.RegisterResp "成功响应"
// @Failure 400 {object} auth.RegisterResp "请求参数错误"
// @Router /auth/register [post]
func (h *RegisterService) Run(req *auth.RegisterReq) (resp *auth.RegisterResp, err error) {
	if req == nil {
		return &auth.RegisterResp{
			StatusCode: -1,
			StatusMsg:  "参数为空",
		}, nil
	}
	if req.ConfirmPassword != req.Password {
		return &auth.RegisterResp{
			StatusCode: -1,
			StatusMsg:  "密码输入不一致",
		}, nil
	}

	userData := &rpcuser.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		Username:        req.Username,
		Telephone:       req.Telephone,
	}
	//创建用户
	_, err = rpc.UserClient.Register(h.Context, userData)
	if err != nil {
		return &auth.RegisterResp{
			StatusCode: -1,
			StatusMsg:  "注册失败",
		}, err
	}

	//绑定角色组
	//userIdStr := fmt.Sprintf("%d", register.UserId)
	//if _, err = mw.Enforcer.AddGroupingPolicy(userIdStr, "user"); err != nil {
	//	klog.Fatal("角色绑定失败:", err)
	//	return &auth.RegisterResp{
	//		StatusCode: -1,
	//		StatusMsg:  "注册失败",
	//	}, err
	//}

	return &auth.RegisterResp{
		StatusCode: 0,
		StatusMsg:  "注册成功",
	}, err
}
