package user

import (
	"context"
	"errors"
	"fmt"
	"go-zero-container/app/usercerter/cmd/models"
	"go-zero-container/app/usercerter/cmd/rpc/usercenter"
	common_models "go-zero-container/common/global/models"
	"gorm.io/gorm"

	"go-zero-container/app/usercerter/cmd/api/internal/svc"
	"go-zero-container/app/usercerter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLogic {
	return &CancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLogic) Cancel(req *types.CancelReq) (resp *types.CancelResp, err error) {
	// todo: add your logic here and delete this line

	// 软删除用户表中对应的用户
	var user models.SysUser
	if errors.Is(l.svcCtx.DB.First(&user, "uuid = ?", req.Uuid).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("被注销的用户不存在")
	}

	// 检查被注销用户是否存在容器
	var container common_models.Container
	err = l.svcCtx.DB.First(&container, "user_uuid = ?", req.Uuid).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("数据库获取用户信息失败")
	} else if err != gorm.ErrRecordNotFound {
		return nil, errors.New("用户存在容器无法注销")
	}

	userName := fmt.Sprintf("%s", l.ctx.Value("userName"))
	// 使用rpc服务注销用户
	res, err := l.svcCtx.UsercenterRpc.Cancel(l.ctx, &usercenter.CancelReq{
		Username: userName,
		Uuid:     req.Uuid,
	})

	if err != nil {
		return nil, errors.New("注销用户失败，请稍后重试")
	}

	return &types.CancelResp{
		Code:    res.Code,
		Message: res.Message,
	}, nil
}
