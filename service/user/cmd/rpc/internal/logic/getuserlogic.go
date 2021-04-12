package logic

import (
	"context"
	"fmt"

	"book/service/user/cmd/rpc/internal/svc"
	"book/service/user/cmd/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	// todo: add your logic here and delete this line
	fmt.Println(222)
	one, err := l.svcCtx.UserModel.FindOne(in.Id)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoReply{
		Id:     one.Id,
		Name:   one.Name,
		Number: one.Number,
		Gender: one.Gender,
	}, nil
	//return &user.UserInfoReply{}, nil
}
