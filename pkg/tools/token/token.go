package token

import (
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/fwhezfwhez/errorx"
	"time"
)

const JWT_SECRETE = "p2F*J7D!A%s68^wS"
const EXPIRE_TIME = time.Hour * 3

func GenerateToken(user model.Staff) string {
	token := jwt.New(jwt.SigningMethodES384)
	claims, now := make(jwt.MapClaims, 0), time.Now()
	claims["userID"] = user.RecID
	claims["userName"] = user.StaffName
	claims["exp"] = now.Add(EXPIRE_TIME)
	token.Claims = claims
	tokenStr, err := token.SignedString([]byte(JWT_SECRETE))
	if err != nil {
		logger.Error(errorx.MustWrap(err), "生成 Token 失败")
	}
	return tokenStr
}

func IsExpired(tokenStr string) bool {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRETE), nil
	})
	if err != nil {
		logger.Error(errorx.MustWrap(err), "解析 Token 失败")
		return true
	}
	if token != nil {
		if _, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {

			//m := map[string]interface{}(claims)

		}
	}
	return false

}
