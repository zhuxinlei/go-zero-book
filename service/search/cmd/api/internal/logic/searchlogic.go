package logic

import (
	"book/service/user/cmd/rpc/userclient"
	"context"
	"encoding/json"
	"fmt"

	"book/service/search/cmd/api/internal/svc"
	"book/service/search/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchReq) (*types.SearchReply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("userId--------",l.ctx.Value("userId"))

	//调用user rpc服务中的getUser方法
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}
	fmt.Println(req.Name)
	// 使用user rpc
	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(userInfo)
	return &types.SearchReply{
		Name:  req.Name,
		Count: 100,
	}, nil
}
