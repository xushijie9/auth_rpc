// Code generated by goctl. DO NOT EDIT.
// Source: auth_rpc.proto

package server

import (
	"context"

	"github.com/xushijie9/auth_rpc/rpc/internal/logic"
	"github.com/xushijie9/auth_rpc/rpc/internal/svc"
	"github.com/xushijie9/auth_rpc/rpc/types/auth_rpc"
)

type AuthServer struct {
	svcCtx *svc.ServiceContext
	auth_rpc.UnimplementedAuthServer
}

func NewAuthServer(svcCtx *svc.ServiceContext) *AuthServer {
	return &AuthServer{
		svcCtx: svcCtx,
	}
}

func (s *AuthServer) LoginAuth(ctx context.Context, in *auth_rpc.AuthRequest) (*auth_rpc.AuthReply, error) {
	l := logic.NewLoginAuthLogic(ctx, s.svcCtx)
	return l.LoginAuth(in)
}
