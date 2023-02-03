package wechat

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// GetFriendAndChatRoomList 获取好友和群清单 https://www.showdoc.com.cn/WeChatProject/8995071288617868
func (w *Wechat) GetFriendAndChatRoomList() (*model.FriendAndChatRoomResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return nil, err
	}

	resultBody, err := client.DoPost("/GetFriendAndChatRoomList", nil)
	if err != nil {
		return nil, err
	}

	//fmt.Println("resultBody:", string(resultBody))

	friendAndChatRoom := model.FriendAndChatRoomResult{}
	err = json.Unmarshal(resultBody, &friendAndChatRoom)
	if err != nil {
		return nil, err
	}

	return &friendAndChatRoom, nil
}

// StickyChat 置顶聊天 https://www.showdoc.com.cn/WeChatProject/9063350563096436
func (w *Wechat) StickyChat(wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.StickyChatReq{
		Gidorwxid: wxid,
	}

	resultBody, err := client.DoPost("/StickyChat", req)
	if err != nil {
		return err
	}

	commonResult := model.StickyChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.StickyChat != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// UnpinChat 取消置顶聊天 https://www.showdoc.com.cn/WeChatProject/9063351366305816
func (w *Wechat) UnpinChat(wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.StickyChatReq{
		Gidorwxid: wxid,
	}

	resultBody, err := client.DoPost("/UnpinChat", req)
	if err != nil {
		return err
	}

	commonResult := model.UnpinChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.UnpinChat != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// BatchGetContactBriefInfo 批量获取联系人简明信息(一次最多传100个id) https://www.showdoc.com.cn/WeChatProject/9333452323089279
func (w *Wechat) BatchGetContactBriefInfo(wxids []string) (*[]model.ContactBriefInfo, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return nil, err
	}

	req := model.BatchGetContactBriefInfoReq{
		WxidList: strings.Join(wxids, ","),
	}

	resultBody, err := client.DoPost("/BatchGetContactBriefInfo", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.BatchGetContactBriefInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult.Info, nil
}

// GetFriendOrChatroomDetailInfo 方法_网络获取好友或群详细信息,好友则返回好友信息，群聊则返回群成员数量。成员微信id，成员昵称 https://www.showdoc.com.cn/WeChatProject/9157629212860224
func (w *Wechat) GetFriendOrChatroomDetailInfo(wxidorgid string) (*model.GetFriendOrChatroomDetailInfoResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return nil, err
	}

	req := model.GetFriendOrChatroomDetailInfoReq{
		WxidOrGid: wxidorgid,
	}

	resultBody, err := client.DoPost("/GetFriendOrChatroomDetailInfo", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.GetFriendOrChatroomDetailInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}

// MarkAsReadSession 标为已读会话 https://www.showdoc.com.cn/WeChatProject/9219474016553735
func (w *Wechat) MarkAsReadSession(gidOrWxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.MarkAsReadSessionReq{
		GidOrWxid: gidOrWxid,
	}

	resultBody, err := client.DoPost("/MarkAsReadSession", req)
	if err != nil {
		return err
	}

	commonResult := model.MarkAsReadSessionResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.MarkAsReadSession != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// InitContact 初始化通讯录 https://www.showdoc.com.cn/WeChatProject/9730605278628502
func (w *Wechat) InitContact(url string) (*[]model.InitContactResultBatchItem, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return nil, err
	}

	resultBody, err := client.DoPost("/InitContact", nil)
	if err != nil {
		return nil, err
	}

	commonResult := model.InitContactResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	if commonResult.InitContact != "1" {
		return nil, fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return &commonResult.Batch, nil
}

// GetCurrentChatObjectInfo 获取当前聊天对象信息 https://www.showdoc.com.cn/WeChatProject/9019586393938696
func (w *Wechat) GetCurrentChatObjectInfo(url string) (*model.GetCurrentChatObjectInfoResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return nil, err
	}

	resultBody, err := client.DoPost("/GetCurrentChatObjectInfo", nil)
	if err != nil {
		return nil, err
	}

	commonResult := model.GetCurrentChatObjectInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}

// SwitchCurrentChatObject 切换当前聊天对象 https://www.showdoc.com.cn/WeChatProject/9019589526757527
func (w *Wechat) SwitchCurrentChatObject(wxidOrGid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.SwitchCurrentChatObjectReq{
		WxidOrGid: wxidOrGid,
	}

	resultBody, err := client.DoPost("/SwitchCurrentChatObject", req)
	if err != nil {
		return err
	}

	commonResult := model.SwitchCurrentChatObjectResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SwitchCurrentChatObject != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// TurnOnDoNotDisturb 开启消息免打扰 https://www.showdoc.com.cn/WeChatProject/9063347648143133
func (w *Wechat) TurnOnDoNotDisturb(gidOrWxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.TurnOnDoNotDisturbReq{
		GidOrWxid: gidOrWxid,
	}

	resultBody, err := client.DoPost("/TurnOnDoNotDisturb", req)
	if err != nil {
		return err
	}

	commonResult := model.TurnOnDoNotDisturbResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.TurnOnDoNotDisturb != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// TurnOffDoNotDisturb 关闭消息免打扰 https://www.showdoc.com.cn/WeChatProject/9063348367251383
func (w *Wechat) TurnOffDoNotDisturb(gidOrWxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.TurnOffDoNotDisturbReq{
		GidOrWxid: gidOrWxid,
	}

	resultBody, err := client.DoPost("/TurnOffDoNotDisturb", req)
	if err != nil {
		return err
	}

	commonResult := model.TurnOffDoNotDisturbResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.TurnOffDoNotDisturb != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SaveToContact 保存到通讯录 https://www.showdoc.com.cn/WeChatProject/9063365453994960
func (w *Wechat) SaveToContact(gidOrWxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.SaveToContactReq{
		GidOrWxid: gidOrWxid,
	}

	resultBody, err := client.DoPost("/SaveToContact", req)
	if err != nil {
		return err
	}

	commonResult := model.SaveToContactResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SaveToContact != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// RemoveToContact 从通讯录移除 https://www.showdoc.com.cn/WeChatProject/9063365712246544
func (w *Wechat) RemoveToContact(gidOrWxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.RemoveToContactReq{
		GidOrWxid: gidOrWxid,
	}

	resultBody, err := client.DoPost("/RemoveToContact", req)
	if err != nil {
		return err
	}

	commonResult := model.RemoveToContactResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.RemoveToContact != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// MarkAsNoReadSession 标为未读会话 https://www.showdoc.com.cn/WeChatProject/9644356701774840
func (w *Wechat) MarkAsNoReadSession(gidOrWxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.MarkAsNoReadSessionReq{
		GidOrWxid: gidOrWxid,
	}

	resultBody, err := client.DoPost("/MarkAsNoReadSession", req)
	if err != nil {
		return err
	}

	commonResult := model.MarkAsNoReadSessionResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.MarkAsNoReadSession != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
