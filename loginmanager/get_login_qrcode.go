package loginmanager

import "encoding/json"

type GetLoginQRCodeResult struct {
	HexData string `json:"hex_data"`
}

// GetLoginQRCode 获取登录二维码图片 https://www.showdoc.com.cn/WeChatProject/9026367939952187
func (l *LoginManagerServiceImpl) GetLoginQRCode() (string, error) {

	resultBody, err := l.http.DoPost("/GetLoginQRCode", nil)
	if err != nil {
		return "", err
	}

	commonResult := GetLoginQRCodeResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.HexData, nil
}
