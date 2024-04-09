package logic

import (
	"context"
	"errors"
	common_models "go-zero-container/common/global/models"
	jwt_token "go-zero-container/common/jwt-token"

	"go-zero-container/app/usercerter/cmd/rpc/internal/svc"
	"go-zero-container/app/usercerter/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line

	var user common_models.SysUser
	err := l.svcCtx.DB.First(&user, "username = ? and password = ?", in.Username, in.Password).Error
	if err != nil {
		return nil, errors.New("用户账号或密码错误")
	}

	// 对账户状态进行判断，1->正常用户 2 ->被封号用户
	if user.Enable != 1 {
		return nil, errors.New("当前账户已被限制登录，请联系管理员咨询解决方案")
	}

	// 加载 jwt 相关配置
	gtr, err := jwt_token.GenerateToken(user)

	if err != nil {
		return nil, errors.New("创建jwt失败，请稍后重试")
	}

	return &pb.LoginResp{
		NickName:     user.NickName,
		HeaderImg:    user.HeaderImg,
		AccessToken:  gtr.AccessSecret,
		AccessExpire: gtr.AccessExpire,
	}, nil
}
