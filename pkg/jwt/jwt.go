package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// MyClaims 注意这里不要 存储 密码之类的敏感信息哟
type MyClaims struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var mySerect = []byte("wuyue is good man")

// GenToken 生成token
func GenToken(username string, userid int64) (string, error) {
	c := MyClaims{
		UserId:   userid,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).UnixNano(), //过期时间
			Issuer:    "bbs-project",                                  //签发人
		},
	}
	// 加密这个token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 用签名来 签名这个token
	return token.SignedString(mySerect)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {

	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySerect, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if token.Valid {
		return mc, nil
	}

	return nil, errors.New("invalid token")

}
