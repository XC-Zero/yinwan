package token

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/encode"
	"log"
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
		return true
	}
	if len(result) == 0 {
		return true
	}
	aes, err := encode.DecryptByAes(tokenStr)
	if err != nil {
		log.Println(err)
		return true
	}
	log.Println(string(aes), staffEmail)
	if string(aes) != staffEmail {
		log.Println(string(aes) == staffEmail)
		return true
	}
	return false

}
