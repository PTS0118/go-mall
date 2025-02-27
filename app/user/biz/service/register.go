package service

import (
	"context"
	utils2 "github.com/PTS0118/go-mall/api/biz/utils"
	"github.com/PTS0118/go-mall/app/user/biz/model"
	user "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 1. 创建用户（数据库操作）
	userData := model.User{
		Username:  req.Username,
		Email:     req.Email,
		Telephone: req.Telephone,
		Password:  utils2.MD5(req.Password),
		Role:      "user", //默认角色
	}
	err = model.Create(s.ctx, &userData)
	if err != nil {
		return &user.RegisterResp{
			UserId: 0,
		}, err
	}

	return &user.RegisterResp{
		UserId: int32(userData.Id),
	}, nil
}
