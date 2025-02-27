package service

import (
	"context"
	"github.com/PTS0118/go-mall/app/user/biz/model"
	user "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	userData, err := model.FindUserByNameOrEmail(&req.Username, &req.Email, req.GetUserId())
	if err != nil {
		return nil, err
	}
	resp = &user.GetUserResp{
		Id:        int32(userData.Base.Id),
		Email:     userData.Email,
		Username:  userData.Username,
		Telephone: userData.Telephone,
	}
	return resp, err
}
