package wechat

import (
	"encoding/json"

	"fmt"
	"strconv"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
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

	commonResult := model.SendTextMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendTextMsg != "1" {
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

// SendGIFMsg 发送GIF动画表情消息 https://www.showdoc.com.cn/WeChatProject/8929127822471500
func SendGIFMsg(url string, wxid, gifPath string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SendGIFMsgReq{
		Wxid:    wxid,
		GifPath: gifPath,
	}

	resultBody, err := client.DoPost("/SendGIFMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendGIFMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendGIFMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendLocationMsg 发送位置消息 https://www.showdoc.com.cn/WeChatProject/9167445492364225
func SendLocationMsg(url string, wxid, msg string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SendLocationMsgReq{
		Wxid: wxid,
		Msg:  msg,
	}

	resultBody, err := client.DoPost("/SendLocationMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendLocationMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendLocationMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendPatMsg 发送拍一拍消息 https://www.showdoc.com.cn/WeChatProject/9644354859714189
func SendPatMsg(url string, wxid, gid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SendPatMsgReq{
		Wxid: wxid,
		Gid:  gid,
	}

	resultBody, err := client.DoPost("/SendLocationMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendPatMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendPatMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendVoiceMsg 发送拍一拍消息 https://www.showdoc.com.cn/WeChatProject/9606012038140554
func SendVoiceMsg(url string, wxid, voiceHex string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SendVoiceMsgReq{
		Wxid:     wxid,
		VoiceHex: voiceHex,
	}

	resultBody, err := client.DoPost("/SendVoiceMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendVoiceMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendVoiceMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendDIYMsg 发送自定义消息 https://www.showdoc.com.cn/WeChatProject/9167460676535402
func SendDIYMsg(url string, wxid, msgType, msg string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SendDIYMsgReq{
		Type: msgType,
		Wxid: wxid,
		Msg:  msg,
	}

	resultBody, err := client.DoPost("/SendDIYMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendDIYMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendDIYMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendCardMsg 转发好友名片 https://www.showdoc.com.cn/WeChatProject/8929459103947241
func SendCardMsg(url string, req model.SendCardMsgReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/SendCardMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendCardMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendCardMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// FowardEssayMsg 转发文章消息 https://www.showdoc.com.cn/WeChatProject/8929539190505167
func FowardEssayMsg(url string, req model.SendCardMsgReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/FowardEssayMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.FowardEssayMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.FowardEssayMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// FowardAppMsg 转发小程序 https://www.showdoc.com.cn/WeChatProject/8929638412121701
func FowardAppMsg(url string, req model.FowardAppMsgReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/FowardAppMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.FowardAppMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.FowardAppMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// FowardMusicMsg 转发网易云音乐 https://www.showdoc.com.cn/WeChatProject/8929642778231025
func FowardMusicMsg(url string, req model.FowardMusicMsgReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/FowardMusicMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.FowardMusicMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.FowardMusicMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// FowardXMLMsg 转发自定义XML消息 https://www.showdoc.com.cn/WeChatProject/9167364405161045
func FowardXMLMsg(url string, req model.FowardMusicMsgReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/FowardXMLMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.FowardXMLMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.FowardXMLMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendAtMsg 发送群@消息 群聊中使用 https://www.showdoc.com.cn/WeChatProject/8929229624237103
func SendAtMsg(url string, req model.SendAtMsgReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/SendAtMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendAtMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendAtMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SendAtAllMsg 发送@所有人 消息 群聊中使用(本人必须是群主或群管理员，否则发送失败) https://www.showdoc.com.cn/WeChatProject/8981408454846254
func SendAtAllMsg(url string, req model.SendAtAllMsgReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/SendAtAllMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.SendAtAllMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendAtAllMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// ForwardAllMsg 转发任意消息（注意：无法转发语音） https://www.showdoc.com.cn/WeChatProject/9090147365509163
func ForwardAllMsg(url string, req model.ForwardAllMsgReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/ForwardAllMsg", req)
	if err != nil {
		return err
	}

	commonResult := model.ForwardAllMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ForwardAllMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// ConfigureMsgRecive 开启/关闭实时消息接收功能。 https://www.showdoc.com.cn/WeChatProject/9204125262722344
func ConfigureMsgRecive(url string, isEnable int, msgReciveFullURL string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.ConfigureMsgReciveReq{
		IsEnable: strconv.Itoa(isEnable),
		Url:      msgReciveFullURL,
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

// GetUnReadMsgNum 获取未读消息总数 https://www.showdoc.com.cn/WeChatProject/9690385454945788
func GetUnReadMsgNum(url string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return "", err
	}

	resultBody, err := client.DoPost("/GetUnReadMsgNum", nil)
	if err != nil {
		return "", err
	}

	commonResult := model.GetUnReadMsgNumResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.GetUnReadMsgNum, nil
}
