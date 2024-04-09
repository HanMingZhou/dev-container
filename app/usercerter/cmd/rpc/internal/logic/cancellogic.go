package logic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/usercerter/cmd/rpc/internal/svc"
	"go-zero-container/app/usercerter/cmd/rpc/pb/pb"
	"go-zero-container/common/models"
)

type CancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLogic {
	return &CancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注销
func (l *CancelLogic) Cancel(in *pb.CancelReq) (*pb.CancelResp, error) {
	// todo: add your logic here and delete this line
	var user models.SysUser
	// 检查用户是否存在，存在则删除，不存在则报错
	err := l.svcCtx.DB.Delete(&user, "uuid = ?", in.Uuid).Error
	if err != nil {
		return nil, errors.New("注销用户失败，请稍后重试")
	}
	return &pb.CancelResp{
		Code:    200,
		Message: "注销成功",
	}, nil

}
