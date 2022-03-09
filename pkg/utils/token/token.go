package token

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/encode"
	"time"
)

//goland:noinspection GoSnakeCaseUsage
const EXPIRE_TIME = time.Second * 3

func GenerateToken(staffEmail string) (string, error) {
	tokenStr, err := encode.EncryptByAes(staffEmail)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// IsExpired token 是否过期
func IsExpired(tokenStr, staffEmail string) bool {
	result, err := client.RedisClient.Get(tokenStr).Result()
	if err != nil {
		return false
	}
	if len(result) == 0 {
		return false
	}
	aes, err := encode.DecryptByAes(tokenStr)
	if err != nil {
		return false
	}
	if string(aes) != staffEmail {
		return false
	}
	return true

}
