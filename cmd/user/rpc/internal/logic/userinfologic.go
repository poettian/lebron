package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhoushuguang/lebron/cmd/user/rpc/internal/svc"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//func (l *UserInfoLogic) AddUser(in *user.AddUserReq) (*user.AddUserRes, error) {
//
//}
//
//func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
//	umsMember, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
//	if err != nil {
//		if err == model.ErrNotFound {
//			return nil, status.Error(100, "用户不存在")
//		}
//		return nil, status.Error(500, err.Error())
//	}
//	var resp user.UserInfo
//	_ = copier.Copy(&resp, umsMember)
//	resp.CreateTime = umsMember.CreateTime.Unix()
//	resp.UpdateTime = umsMember.UpdateTime.Unix()
//	return &user.UserInfoResponse{
//		User: &resp,
//	}, nil
//}
