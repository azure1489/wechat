package httptools

import (
	"encoding/base64"
	"encoding/json"

	"github.com/azure1489/wechat/util/encryptedtools"
)

func GetSendRequestBody(proxyBody ProxyBody, publicKeyPath string) ([]byte, error) {

	original, err := json.Marshal(proxyBody)
	if err != nil {
		return nil, err
	}

	encryptedAESKey, encryptedData, err := encryptedtools.EncryptedData(original, publicKeyPath)
	if err != nil {
		return nil, err
	}

	// 准备请求体
	requestBody, err := json.Marshal(ProxyReq{
		EncryptedAESKey: base64.StdEncoding.EncodeToString(encryptedAESKey),
		EncryptedData:   base64.StdEncoding.EncodeToString(encryptedData),
	})
	if err != nil {
		return nil, err
	}

	return requestBody, nil

}
