package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func encryptConfigFile(originFilePath, encryptFilePath, key string) (err error) {

	// 读加密源文件
	plainText, err := os.ReadFile(originFilePath)
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
		return
	}

	// 算法块
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
		return
	}

	// 在 CTR 加密的基础上增加 GMAC 的特性，解决了 CTR 不能对加密消息进行完整性校验的问题。
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
		return
	}

	// 随机数
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
		return
	}

	// 加密：明文->密文
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	// 写密文到文件
	err = os.WriteFile(encryptFilePath, cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
		return
	}
	return
}
