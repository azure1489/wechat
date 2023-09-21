package loginmanager

// RefreshLoginQRCode 刷新登录二维码 https://www.showdoc.com.cn/WeChatProject/8966162223712985
func (l *LoginManagerServiceImpl) RefreshLoginQRCode(url string) ([]byte, error) {

	resultBody, err := l.http.DoGet("/RefreshLoginQRCode")
	if err != nil {
		return nil, err
	}

	return resultBody, nil
}
