package service

import (
	"context"
	user "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	////klog.Info("LoginRqp:%+v", req)
	//if mysql.DB != nil {
	//	userLogin, _ := model.GetByEmail(mysql.DB, s.ctx, "test@test")
	//	fmt.Printf("user%+v", userLogin)
	//} else {
	//	println("mysql is nil")
	//}
	////if err != nil {
	////	klog.Error("LoginError:%+v", err)
	////}
	return
}
