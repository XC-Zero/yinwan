package encode

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
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
	tail := len(plainByte) % aes.BlockSize
	if tail != 0 {
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
  AES  GCM 解码
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

//goland:noinspection GoSnakeCaseUsage
const JWT_SECRETE = "p2F*J7D!A%s68^wS"

var pwdKey = []byte(JWT_SECRETE)

//pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

//pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

//AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	encrypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return encrypted, nil
}

//AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	encrypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(encrypted, data)
	//去除填充
	encrypted, err = pkcs7UnPadding(encrypted)
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

//EncryptByAes Aes加密 后 base64 再加
func EncryptByAes(text string) (string, error) {
	res, err := AesEncrypt([]byte(text), pwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

//DecryptByAes Aes 解密
func DecryptByAes(data string) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return AesDecrypt(dataByte, pwdKey)
}
