package logic

import (
	"context"
	"errors"

	"auth_rpc/rpc/internal/svc"
	"auth_rpc/rpc/types/auth_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginAuthLogic {
	return &LoginAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginAuthLogic) LoginAuth(in *auth_rpc.AuthRequest) (*auth_rpc.AuthReply, error) {
	// todo: add your logic here and delete this line
	if in.GetUserName() == "admin" && in.GetUserPass() == "youPass" {
		return &auth_rpc.AuthReply{
			UserName:       in.GetUserName(),
			Msg:            "login successful",
			IsLoginSuccess: true,
		}, nil
	}
	return nil, errors.New("user login is failed")
}
