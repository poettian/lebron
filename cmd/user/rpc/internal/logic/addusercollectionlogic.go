package logic

import (
	"context"

	"github.com/zhoushuguang/lebron/cmd/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/cmd/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserCollectionLogic {
	return &AddUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加收藏
func (l *AddUserCollectionLogic) AddUserCollection(in *user.UserCollectionAddReq) (*user.UserCollectionAddRes, error) {
	// todo: add your logic here and delete this line

	return &user.UserCollectionAddRes{}, nil
}
