package wechat

import (
	"encoding/json"

	"fmt"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// AgainStartWeChat 启动更多微信 https://www.showdoc.com.cn/WeChatProject/9063540299207309
func (w *Wechat) AgainStartWeChat(req model.AgainStartWeChatReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/AgainStartWeChat", req)
	if err != nil {
		return err
	}

	commonResult := model.AgainStartWeChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.AgainStartWeChat != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// RefreshLoginQRCode 刷新登录二维码 https://www.showdoc.com.cn/WeChatProject/8966162223712985
func (w *Wechat) RefreshLoginQRCode(url string) ([]byte, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return nil, err
	}

	resultBody, err := client.DoGet("/RefreshLoginQRCode")
	if err != nil {
		return nil, err
	}

	return resultBody, nil
}

// GetLoginQRCode 获取登录二维码图片 https://www.showdoc.com.cn/WeChatProject/9026367939952187
func (w *Wechat) GetLoginQRCode(url string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return "", err
	}

	resultBody, err := client.DoPost("/GetLoginQRCode", nil)
	if err != nil {
		return "", err
	}

	commonResult := model.GetLoginQRCodeResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.HexData, nil
}

// IsLoginStatus 获取微信登陆状态 https://www.showdoc.com.cn/WeChatProject/8991915897471323
func (w *Wechat) IsLoginStatus(url string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/IsLoginStatus", nil)
	if err != nil {
		return err
	}

	commonResult := model.IsLoginStatusResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.OnlineStatus != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// GetSelfLoginInfo 获取个人详细信息 https://www.showdoc.com.cn/WeChatProject/8929111706614173
func (w *Wechat) GetSelfLoginInfo(url string) (*model.GetSelfLoginInfoResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return nil, err
	}

	resultBody, err := client.DoPost("/GetSelfLoginInfo", nil)
	if err != nil {
		return nil, err
	}

	commonResult := model.GetSelfLoginInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}

// Logout 退出微信 https://www.showdoc.com.cn/WeChatProject/9008665345790557
func (w *Wechat) Logout(url string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/Logout", nil)
	if err != nil {
		return err
	}

	commonResult := model.LogoutResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// TerminateThisWeChat 结束微信 https://www.showdoc.com.cn/WeChatProject/9214210657048561
func (w *Wechat) TerminateThisWeChat(url string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/TerminateThisWeChat", nil)
	if err != nil {
		return err
	}

	commonResult := model.TerminateThisWeChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
