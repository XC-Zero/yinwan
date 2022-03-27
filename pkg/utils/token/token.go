package token

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/encode"
	"log"
	"strings"
	"time"
)

//goland:noinspection GoSnakeCaseUsage
const EXPIRE_TIME = time.Second * 3
const SPLIT_SYMBOL = "|||||"

func GenerateToken(staffEmail string) (string, error) {
	str := staffEmail + SPLIT_SYMBOL + time.Now().String()
	log.Printf(str)
	tokenStr, err := encode.EncryptByAes(str)
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
	email := strings.Split(string(aes), SPLIT_SYMBOL)[0]
	log.Printf("aes email is %s \n input email is %s ", email, staffEmail)
	if email != staffEmail {
		return true
	}
	return false

}
