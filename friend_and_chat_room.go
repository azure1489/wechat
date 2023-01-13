package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
	"strings"
	"time"
)

// GetFriendAndChatRoomList 获取好友和群清单 https://www.showdoc.com.cn/WeChatProject/8995071288617868
func GetFriendAndChatRoomList(url string) (*model.FriendAndChatRoomResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
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
func StickyChat(url string, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
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
func UnpinChat(url string, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
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
func BatchGetContactBriefInfo(url string, wxids []string) (*[]model.ContactBriefInfo, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
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
func GetFriendOrChatroomDetailInfo(url string, wxidorgid string) (*model.GetFriendOrChatroomDetailInfoResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
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
func MarkAsReadSession(url string, gidOrWxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
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
