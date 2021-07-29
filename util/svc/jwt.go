package svc

import (
	"encoding/json"
	jwtBase "github.com/golang-jwt/jwt"
	"github.com/wuzehv/ginframework/model"
	"github.com/wuzehv/ginframework/service/db"
	"github.com/wuzehv/ginframework/util/config"
	"github.com/wuzehv/ginframework/util/journal"
	"github.com/wuzehv/ginframework/util/jwt"
	"github.com/wuzehv/ginframework/util/static"
)

type j struct {
	*model.User
}

func (_ *j) Generate(userId, _ uint) (string, error) {
	var u model.User
	if err := db.Db.First(&u, userId).Error; err != nil {
		return "", err
	}

	return jwt.GenToken(j{&u}, "", config.Svc.ExpireTime)
}

func (_ *j) Confirm(token string) error {
	_, err := validToken(token)
	return err
}

func (_ *j) Valid(token string, user *model.User) error {
	x, err := validToken(token)

	if err != nil {
		return err
	}

	// 解析出用户信息
	c, err := json.Marshal(x.Data)
	if err != nil {
		return err
	}

	return json.Unmarshal(c, user)
}

func validToken(token string) (*jwt.Claims, error) {
	x, err := jwt.ValidToken(token, "")

	if err != nil {
		switch err.(*jwtBase.ValidationError).Errors {
		case jwtBase.ValidationErrorExpired:
			return nil, static.SessionExpired
		default:
			journal.Error("jwt_svc", err)
			return nil, static.SystemError
		}
	}

	return x, nil
}

func (_ *j) Destroy(_ uint) error {
	return nil
}
