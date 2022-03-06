package token

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/encode"
	"time"
)

//goland:noinspection GoSnakeCaseUsage
const EXPIRE_TIME = time.Second * 3

func GenerateToken(userID string) (string, error) {
	tokenStr, err := encode.EncryptByAes(userID)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
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
