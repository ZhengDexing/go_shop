package security

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go_shop/entity"
	"time"
)

type Jwt struct{}

// 创建token
func (j Jwt) CreateToken(user entity.User) (string, error) {
	//自定义claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"nickName": user.NickName,
		"exp":      100,
		// 生效时间
		"nbf": time.Now().Unix(),
		// 签发时间
		"iat": time.Now().Unix(),
	})
	return token.SignedString([]byte(""))
}

// 验证token
func (j Jwt) ParseToken(tokens string) (*entity.User, error) {
	user := new(entity.User)
	token, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return nil, err
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return nil, err
	}

	user.Id = claim["id"].(string)
	user.NickName = claim["nickName"].(string)
	return user, nil
}
