package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
)

const (
	SecretKey = "Ihavelogin"
)

//获取解密后的token
func ParseToken(s string) (*jwt.Token, error) {
	fn := func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	}
	return jwt.Parse(s, fn)
}

//创建token
func CreateToken(id,mobile, lastLogin string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"mobile": mobile,
		"last_login":lastLogin,
	})
	if tokenString, err := token.SignedString([]byte(SecretKey)); err == nil {
		return tokenString
	} else {
		return ""
	}
}

//获取解密的参数
func GetFromClaims(claims jwt.Claims) map[string]string {
	result := make(map[string]string)
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			result[fmt.Sprintf("%s", k.Interface())] = fmt.Sprintf("%v", value.Interface())
		}
	}
	return result
}