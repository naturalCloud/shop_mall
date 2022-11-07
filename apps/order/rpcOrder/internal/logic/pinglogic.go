package logic

import (
	"context"

	"shop_mall/apps/order/rpcOrder/internal/svc"
	"shop_mall/apps/order/rpcOrder/rpcOrder"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *rpcOrder.Request) (*rpcOrder.Response, error) {
	// todo: add your logic here and delete this line

	return &rpcOrder.Response{}, nil
}
