// Code generated by goctl. DO NOT EDIT!
// Source: pay.proto

package server

import (
	"context"

	"go-zero-admin/rpc/pay/internal/logic"
	"go-zero-admin/rpc/pay/internal/svc"
	"go-zero-admin/rpc/pay/pay"
)

type PayServer struct {
	svcCtx *svc.ServiceContext
}

func NewPayServer(svcCtx *svc.ServiceContext) *PayServer {
	return &PayServer{
		svcCtx: svcCtx,
	}
}

func (s *PayServer) Ping(ctx context.Context, in *pay.Request) (*pay.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
