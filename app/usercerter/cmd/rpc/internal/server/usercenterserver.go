// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package server

import (
	"context"

	"go-zero-container/app/usercerter/cmd/rpc/internal/logic"
	"go-zero-container/app/usercerter/cmd/rpc/internal/svc"
	"go-zero-container/app/usercerter/cmd/rpc/pb/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

// 用户注销
func (s *UsercenterServer) Cancel(ctx context.Context, in *pb.CancelReq) (*pb.CancelResp, error) {
	l := logic.NewCancelLogic(ctx, s.svcCtx)
	return l.Cancel(in)
}

func (s *UsercenterServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UsercenterServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
