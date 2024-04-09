package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/usercerter/cmd/api/internal/svc"
	"go-zero-container/app/usercerter/cmd/api/internal/types"
	"go-zero-container/app/usercerter/cmd/rpc/usercenter"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type GenTokenResq struct {
	AccessToken  string
	AccessExpire int64
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	//var user common_models.SysUser
	//err = l.svcCtx.DB.First(&user, "username = ? and password = ?", req.Username, req.Password).Error
	//if err != nil {
	//	return nil, errors.New("用户账号或密码错误")
	//}
	//
	//// 对账户状态进行判断，1->正常用户 2 ->被封号用户
	//if user.Enable != 1 {
	//	return nil, errors.New("当前账户已被限制登录，请联系管理员咨询解决方案")
	//}
	//
	//// 加载 jwt 相关配置
	//gtr, err := jwt_token.GenerateToken(user)
	//
	//if err != nil {
	//	return nil, errors.New("创建jwt失败，请稍后重试")
	//}
	//
	//return &types.LoginResp{
	//	NickName:     user.NickName,
	//	HeaderImg:    user.HeaderImg,
	//	AccessToken:  gtr.AccessSecret,
	//	AccessExpire: gtr.AccessExpire,
	//}, err
	res, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken:  res.AccessToken,
		AccessExpire: res.AccessExpire,
		NickName:     res.NickName,
		HeaderImg:    res.HeaderImg,
	}, nil

}

//func (l *LoginLogic) GenerateToken(user common_models.SysUser) (gtr *GenTokenResq, err error) {
//	now := time.Now().Unix()
//	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
//	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, user)
//	if err != nil {
//		return nil, errors.New("生成jwt失败")
//	}
//	return &GenTokenResq{
//		AccessToken:  accessToken,
//		AccessExpire: accessExpire,
//	}, nil
//}
//
//func (l *LoginLogic) getJwtToken(secretKey string, iat int64, accessExpire int64, user common_models.SysUser) (string, error) {
//
//	claims := make(jwt.MapClaims)
//	claims["exp"] = iat + accessExpire
//	claims["iat"] = iat
//	claims["ID"] = user.ID
//	claims["UUID"] = user.UUID
//	claims["Nickname"] = user.NickName
//	claims["Username"] = user.Username
//	claims["Authorityid"] = user.AuthorityId
//	token := jwt.New(jwt.SigningMethodHS256)
//	token.Claims = claims
//	return token.SignedString([]byte(secretKey))
//}
