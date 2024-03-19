package encryptedtools

import (
	"crypto/rand"
	"log"
	"os"

	"github.com/azure1489/wechat/util/cryptotools"
)

func EncryptedData(original []byte, publicKeyPath string) ([]byte, []byte, error) {

	// 加载公钥
	publicKeyStr, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, nil, err
	}
	publicKey, err := cryptotools.ParseRsaPublicKeyFromPemStr(string(publicKeyStr))
	if err != nil {
		return nil, nil, err
	}

	// 压缩数据
	compressedData, err := cryptotools.CompressData(original)
	if err != nil {
		return nil, nil, err
	}

	log.Printf("原始数据大小: %d bytes", len(original))
	log.Printf("压缩后数据大小: %d bytes", len(compressedData))

	// 生成AES密钥
	aesKey := make([]byte, 32) // 使用256位密钥
	_, err = rand.Read(aesKey)
	if err != nil {
		return nil, nil, err
	}

	encryptedData, err := cryptotools.EncryptWithAES(aesKey, compressedData)
	if err != nil {
		return nil, nil, err
	}

	// 使用公钥加密AES密钥
	encryptedAESKey, err := cryptotools.EncryptWithPublicKey(aesKey, publicKey)
	if err != nil {
		return nil, nil, err
	}

	return encryptedAESKey, encryptedData, nil
}
