package logic

import (
	"context"

	"github.com/zhoushuguang/lebron/cmd/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/cmd/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCollectionListLogic {
	return &GetUserCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏列表
func (l *GetUserCollectionListLogic) GetUserCollectionList(in *user.UserCollectionListReq) (*user.UserCollectionListRes, error) {
	// todo: add your logic here and delete this line

	return &user.UserCollectionListRes{}, nil
}
