package service

import (
	"context"
	"github.com/PTS0118/go-mall/api/infra/rpc"

	auth "github.com/PTS0118/go-mall/api/hertz_gen/api/auth"
	common "github.com/PTS0118/go-mall/api/hertz_gen/frontend/common"
	rpcuser "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *common.Empty, err error) {

	_, err = rpc.UserClient.Login(h.Context, &rpcuser.LoginReq{Email: req.Email, Password: req.Password})
	//println("hertz:%+v", res)
	if err != nil {
		return
	}
	return
}
