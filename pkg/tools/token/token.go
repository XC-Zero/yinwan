package token

import (
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fwhezfwhez/errorx"
	"github.com/golang-jwt/jwt"
	"time"
)

const JWT_SECRETE = "p2F*J7D!A%s68^wS"
const EXPIRE_TIME = time.Second * 3

func GenerateToken(userID string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, now := make(jwt.MapClaims, 0), time.Now()
	claims["aud"] = userID
	claims["iss"] = "yinwan"
	claims["iat"] = now
	claims["exp"] = now.Add(EXPIRE_TIME)
	token.Claims = claims
	tokenStr, err := token.SignedString([]byte(JWT_SECRETE))
	if err != nil {
		logger.Error(errorx.MustWrap(err), "生成 Token 失败")
	}
	return tokenStr
}

// IsExpired todo 这玩意总是会有panic
func IsExpired(tokenStr string) bool {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRETE), nil
	})
	if err != nil {
		panic(err)
		//logger.Error(errorx.MustWrap(err), "解析 Token 失败")
		return true
	}
	if token != nil {
		if c, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			return c.VerifyExpiresAt(time.Now().Unix(), false)
		}
	}
	return true

}
