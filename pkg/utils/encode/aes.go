package encode

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"io"
)

/*AESCBCEncrypt
  AES  CBC 加密
  key:加密key
  plaintext：加密明文
  ciphertext:解密返回字节字符串[ 整型以十六进制方式显示]

*/
func AESCBCEncrypt(plaintext, key string) (string, error) {
	plainByte := []byte(plaintext)
	keyByte := []byte(key)
	if len(plainByte)%aes.BlockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	cipherByte := make([]byte, aes.BlockSize+len(plainByte))
	iv := cipherByte[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherByte[aes.BlockSize:], plainByte)

	ciphertext := fmt.Sprintf("%x\n", cipherByte)
	return ciphertext, nil
}

/*AESCBCDecrypted
  AES  CBC 解码
  key:解密key
  ciphertext:加密返回的串
  plaintext：解密后的字符串
*/
func AESCBCDecrypted(ciphertext, key string) (string, error) {
	cipherByte, _ := hex.DecodeString(ciphertext)
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	if len(cipherByte) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := cipherByte[:aes.BlockSize]
	cipherByte = cipherByte[aes.BlockSize:]
	if len(cipherByte)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")

	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherByte, cipherByte)

	plaintext := string(cipherByte[:])
	return plaintext, nil
}

/*AESGCMEncrypt
  AES  GCM 加密
  key:加密key
  plaintext：加密明文
  ciphertext:解密返回字节字符串[ 整型以十六进制方式显示]

*/
func AESGCMEncrypt(plaintext, key string) (string, string, error) {
	plainByte := []byte(plaintext)
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", "", err
	}

	// 由于存在重复的风险，请勿使用给定密钥使用超过2^32个随机值。
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	cipherByte := aesGcm.Seal(nil, nonce, plainByte, nil)
	ciphertext := fmt.Sprintf("%x\n", cipherByte)
	nonceText := fmt.Sprintf("%x\n", nonce)
	return ciphertext, nonceText, nil
}

/*AESGCMDecrypted
  AES  CBC 解码
  key:解密key
  ciphertext:加密返回的串
  plaintext：解密后的字符串
*/
func AESGCMDecrypted(ciphertext, key, nonceText string) (string, error) {
	cipherByte, _ := hex.DecodeString(ciphertext)
	nonce, _ := hex.DecodeString(nonceText)
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plainByte, err := aesGcm.Open(nil, nonce, cipherByte, nil)
	if err != nil {
		return "", err
	}
	plaintext := string(plainByte[:])
	return plaintext, nil
}
