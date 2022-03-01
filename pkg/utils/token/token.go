package token

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/encode"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fwhezfwhez/errorx"
	"time"
)

//goland:noinspection GoSnakeCaseUsage
const JWT_SECRETE = "p2F*J7D!A%s68^wS"

//goland:noinspection GoSnakeCaseUsage
const EXPIRE_TIME = time.Second * 3

func GenerateToken(userID string) string {
	tokenStr, err := encode.AESCBCEncrypt(userID, JWT_SECRETE)
	if err != nil {
		return ""
	}
	err = client.RedisClient.Set(tokenStr, nil, EXPIRE_TIME).Err()
	if err != nil {
		logger.Error(errorx.MustWrap(err), "Redis 设置 token 失败! ")
		return ""
	}

	return tokenStr
}

// IsExpired token 是否过期
func IsExpired(tokenStr string) bool {
	result, err := client.RedisClient.Get(tokenStr).Result()
	if err != nil {
		return false
	}
	if len(result) == 0 {
		return false
	}
	return true

}
