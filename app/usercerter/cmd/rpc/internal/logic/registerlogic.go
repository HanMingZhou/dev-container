package logic

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"go-zero-container/app/usercerter/cmd/models"
	"gorm.io/gorm"

	"go-zero-container/app/usercerter/cmd/rpc/internal/svc"
	"go-zero-container/app/usercerter/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// todo: add your logic here and delete this line

	var user models.SysUser

	// 判断用户名是否已存在
	if !errors.Is(l.svcCtx.DB.First(&user, "username = ?", in.Username).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("当前用户名已存在")
	}

	// 用户角色合法性判断
	var sysAuthority models.SysAuthority
	if errors.Is(l.svcCtx.DB.First(&sysAuthority, "authority_id = ?", in.AuthorityId).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户角色不存在")
	}

	// 附加uuid
	user.UUID = uuid.NewV4()

	user.Username = in.Username

	// 密码hash加密
	//user.Password = utils.BcryptHash(req.Password)
	user.Password = in.Password
	user.NickName = in.NickName
	user.HeaderImg = in.HeaderImg
	user.AuthorityId = uint(in.AuthorityId)
	user.Enable = 1 // 1代表正常，2代表被封号
	user.Phone = in.Phone
	user.Email = in.Email

	ok := l.svcCtx.DB.Create(&user)
	if ok.Error != nil {
		return nil, errors.New("用户创建失败")
	}
	return &pb.RegisterResp{
		Code:    200,
		Message: "ok",
	}, nil
}
