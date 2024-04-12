package jwt_token

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	common_models "go-zero-container/common/global/models"
	"time"
)

type GenTokenResq struct {
	AccessSecret string
	AccessExpire int64
}

func NewToken() *GenTokenResq {
	return &GenTokenResq{
		AccessSecret: "c02dkk3f-094y-59d5-bi98-f1s543daca3d",
		AccessExpire: 18000000,
	}
}
func GenerateToken(user common_models.SysUser) (gtr *GenTokenResq, err error) {
	token := NewToken()
	now := time.Now().Unix()
	accessExpire := token.AccessExpire
	accessToken, err := getJwtToken(token.AccessSecret, now, accessExpire, user)
	if err != nil {
		return nil, errors.New("生成jwt失败")
	}
	return &GenTokenResq{
		AccessSecret: accessToken,
		AccessExpire: accessExpire,
	}, nil
}

func getJwtToken(secretKey string, iat int64, accessExpire int64, user common_models.SysUser) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + accessExpire
	claims["iat"] = iat
	claims["ID"] = user.ID
	claims["UUID"] = user.UUID
	claims["Nickname"] = user.NickName
	claims["Username"] = user.Username
	claims["Authorityid"] = user.AuthorityId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
