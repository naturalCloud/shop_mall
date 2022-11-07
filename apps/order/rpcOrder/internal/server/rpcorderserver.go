// Code generated by goctl. DO NOT EDIT!
// Source: rpcOrder.proto

package server

import (
	"context"

	"shop_mall/apps/order/rpcOrder/internal/logic"
	"shop_mall/apps/order/rpcOrder/internal/svc"
	"shop_mall/apps/order/rpcOrder/rpcOrder"
)

type RpcOrderServer struct {
	svcCtx *svc.ServiceContext
	rpcOrder.UnimplementedRpcOrderServer
}

func NewRpcOrderServer(svcCtx *svc.ServiceContext) *RpcOrderServer {
	return &RpcOrderServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcOrderServer) Ping(ctx context.Context, in *rpcOrder.Request) (*rpcOrder.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
