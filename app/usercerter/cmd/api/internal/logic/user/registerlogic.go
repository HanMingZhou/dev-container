package user

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"go-zero-container/app/usercerter/cmd/models"
	"gorm.io/gorm"

	"go-zero-container/app/usercerter/cmd/api/internal/svc"
	"go-zero-container/app/usercerter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	var user models.SysUser

	// 判断用户名是否已存在
	if !errors.Is(l.svcCtx.DB.First(&user, "username = ?", req.Username).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("当前用户名已存在")
	}

	// 用户角色合法性判断
	var sysAuthority models.SysAuthority
	if errors.Is(l.svcCtx.DB.First(&sysAuthority, "authority_id = ?", req.AuthorityId).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户角色不存在")
	}

	// 附加uuid
	user.UUID = uuid.NewV4()

	user.Username = req.Username

	// 密码hash加密
	//user.Password = utils.BcryptHash(req.Password)
	user.Password = req.Password
	user.NickName = req.NickName
	user.HeaderImg = req.HeaderImg
	user.AuthorityId = req.AuthorityId
	user.Enable = 1 // 1代表正常，2代表被封号
	user.Phone = req.Phone
	user.Email = req.Email

	ok := l.svcCtx.DB.Create(&user)
	if ok.Error != nil {
		return nil, errors.New("用户创建失败")
	}
	return &types.RegisterResp{
		Code:    200,
		Message: "ok",
	}, nil

}
