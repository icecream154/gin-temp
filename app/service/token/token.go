package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/my_jwt"
	"time"
)

func CreateAccTokenFactory() *accToken {
	return &accToken{
		userJwt: my_jwt.CreateMyJWT(variable.ConfigYml.GetString("Token.JwtTokenSignKey")),
	}
}

type accToken struct {
	userJwt *my_jwt.JwtSign
}

// 生成token
func (u *accToken) GenerateToken(id int64, account string, status int64,
	expireAt int64) (tokens string, err error) {

	// 自定义token需要包含的参数
	accClaims := my_jwt.AccClaims{
		Id:      id,
		Account: account,
		Status:  status,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,       // 生效开始时间
			ExpiresAt: time.Now().Unix() + expireAt, // 失效截止时间
		},
	}
	return u.userJwt.CreateAccToken(accClaims)
}

// 用户login成功，记录用户token
func (u *accToken) RecordLoginToken(userToken, clientIp string) bool {
	if _, err := u.userJwt.ParseToken(userToken); err == nil {
		return true
	} else {
		return false
	}
}

// 判断token是否未过期
func (u *accToken) isNotExpired(token string) (*my_jwt.AccClaims, int) {
	if accClaims, err := u.userJwt.ParseToken(token); err == nil {
		if time.Now().Unix()-accClaims.ExpiresAt < 0 {
			// token有效
			return accClaims, consts.JwtTokenOK
		} else {
			// 过期的token
			return accClaims, consts.JwtTokenExpired
		}
	} else {
		// 无效的token
		return nil, consts.JwtTokenInvalid
	}
}

// 判断token是否有效（未过期）
func (u *accToken) IsEffective(token string) bool {
	_, code := u.isNotExpired(token)
	if consts.JwtTokenOK == code {
		return true
	}
	return false
}

// 将 token 解析为绑定时传递的参数
func (u *accToken) ParseToken(tokenStr string) (accClaims my_jwt.AccClaims, err error) {
	if accClaims, err := u.userJwt.ParseToken(tokenStr); err == nil {
		return *accClaims, nil
	} else {
		return my_jwt.AccClaims{}, errors.New(my_errors.ErrorsParseTokenFail)
	}
}
