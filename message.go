package wechat

import (
	"encoding/json"

	"fmt"
	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
	"strconv"
	"time"
)

// SendTextMsg 发送文本消息 https://www.showdoc.com.cn/WeChatProject/8929112442643628
func SendTextMsg(url string, wxid, msg string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SendTextMsgReq{
		Wxid: wxid,
		Msg:  msg,
	}

	resultBody, err := client.DoPost("/SendTextMsg", req)
	if err != nil {
		return err
	}

	commonResult := util.CommonResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendPicMsg 发送图片消息 https://www.showdoc.com.cn/WeChatProject/8929124214624404
func SendPicMsg(url string, wxid, picPath, diyFilename string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SendPicMsgReq{
		Wxid:        wxid,
		PicPath:     picPath,
		DiyFilename: diyFilename,
	}

	resultBody, err := client.DoPost("/SendPicMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendPicMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendPicMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendFileMsg 发送文件消息 https://www.showdoc.com.cn/WeChatProject/8929125290388797
func SendFileMsg(url string, wxid, filepath, diyFilename string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SendFileMsgReq{
		Wxid:        wxid,
		Filepath:    filepath,
		DiyFilename: diyFilename,
	}

	resultBody, err := client.DoPost("/SendFileMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendFileMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendFileMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendVideoMsg 发送视频消息 https://www.showdoc.com.cn/WeChatProject/8929126402335432
func SendVideoMsg(url string, wxid, videoPath, diyFilename string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SendVideoMsgReq{
		Wxid:        wxid,
		VideoPath:   videoPath,
		DiyFilename: diyFilename,
	}

	resultBody, err := client.DoPost("/SendVideoMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendVideoMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendVideoMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// ConfigureMsgReciveFullURL 实时消息接收自定义完整URL。配置后微信接收到的消息会发送到指定的完整URL地址。 https://www.showdoc.com.cn/WeChatProject/8992619194450930
func ConfigureMsgReciveFullURL(url string, msgReciveFullURL string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.ConfigureMsgReciveFullURLReq{
		Url: msgReciveFullURL,
	}

	resultBody, err := client.DoPost("/ConfigureMsgReciveFullURL", req)
	if err != nil {
		return err
	}

	commonResult := model.ConfigureMsgReciveFullURLResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ConfigureMsgReciveFullURL != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// ConfigureMsgRecive 开启/关闭实时消息接收功能。 https://www.showdoc.com.cn/WeChatProject/9204125262722344
func ConfigureMsgRecive(url string, isEnable int) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.ConfigureMsgReciveReq{
		IsEnable: strconv.Itoa(isEnable),
	}

	resultBody, err := client.DoPost("/ConfigureMsgRecive", req)
	if err != nil {
		return err
	}

	commonResult := model.ConfigureMsgReciveResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ConfigureMsgRecive != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// CheckFriendStatus 免打扰检测僵尸粉。 https://www.showdoc.com.cn/WeChatProject/9063410601712380
func CheckFriendStatus(url string, wxid string) (int, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return 0, err
	}

	req := model.CheckFriendStatusReq{
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/CheckFriendStatus", req)
	if err != nil {
		return 0, err
	}

	commonResult := model.CheckFriendStatusResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return 0, err
	}

	checkFriendStatus, err := strconv.Atoi(commonResult.CheckFriendStatus)
	if err != nil {
		return 0, err
	}

	return checkFriendStatus, nil
}

// GetGIFURL 获取GIF的访问URL https://www.showdoc.com.cn/WeChatProject/9690133784680867
func GetGIFURL(url string, msgXml string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return "", err
	}

	req := model.GetGIFURLReq{
		MsgXml: msgXml,
	}

	resultBody, err := client.DoPost("/GetGIFURL", req)
	if err != nil {
		return "", err
	}

	commonResult := model.GetGIFURLResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.GetGIFURL, nil
}

// DownPic 下载图片 https://www.showdoc.com.cn/WeChatProject/9682774769136221
func DownPic(url string, toPath, msgXml string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.DownPicReq{
		ToPath: toPath,
		MsgXml: msgXml,
	}

	resultBody, err := client.DoPost("/DownPic", req)
	if err != nil {
		return err
	}

	commonResult := model.DownPicResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.DownPic != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
