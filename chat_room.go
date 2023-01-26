package wechat

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// CreateChatRoom 创建群聊 https://www.showdoc.com.cn/WeChatProject/9019609319630104
func CreateChatRoom(url string, wxidList []string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return "", err
	}

	req := model.CreateChatRoomReq{
		WxidList: strings.Join(wxidList, ","),
	}

	resultBody, err := client.DoPost("/CreateChatRoom", req)
	if err != nil {
		return "", err
	}

	commonResult := model.CreateChatRoomResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	if commonResult.Retstr != "Everything is OK" {
		return "", fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return commonResult.Gid, nil
}

// AddFriendToChatRoom 添加好友进群(40人以下) https://www.showdoc.com.cn/WeChatProject/9020971162414036
func AddFriendToChatRoom(url string, gid string, wxidList []string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.AddFriendToChatRoomReq{
		Gid:      gid,
		WxidList: strings.Join(wxidList, ","),
	}

	resultBody, err := client.DoPost("/AddFriendToChatRoom", req)
	if err != nil {
		return err
	}

	commonResult := model.AddFriendToChatRoomResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.AddFriendToChatRoom != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// InviteFriendToChatRoom 邀请好友进群(40人以上) https://www.showdoc.com.cn/WeChatProject/9021493538921140
func InviteFriendToChatRoom(url string, gid, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.InviteFriendToChatRoomReq{
		Gid:  gid,
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/InviteFriendToChatRoom", req)
	if err != nil {
		return err
	}

	commonResult := model.InviteFriendToChatRoomResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.InviteFriendToChatRoom != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// AgreeJoinGroup 同意加入群聊 https://www.showdoc.com.cn/WeChatProject/9032882631445023
func AgreeJoinGroup(url string, groupAccessUrl string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.AgreeJoinGroupRoomReq{
		GroupAccessUrl: groupAccessUrl,
	}

	resultBody, err := client.DoPost("/AgreeJoinGroup", req)
	if err != nil {
		return err
	}

	commonResult := model.AgreeJoinGroupResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.AgreeJoinGroup != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// ChatRoomEditMark 修改群聊备注 https://www.showdoc.com.cn/WeChatProject/8994766712410900
func ChatRoomEditMark(url string, gid, mark string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.ChatRoomEditMarkReq{
		Gid:  gid,
		Mark: mark,
	}

	resultBody, err := client.DoPost("/ChatRoomEditMark", req)
	if err != nil {
		return err
	}

	commonResult := model.ChatRoomEditMarkResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ChatRoomEditMark != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// ChatRoomEditName 修改群聊名称 https://www.showdoc.com.cn/WeChatProject/8993745484621751
func ChatRoomEditName(url string, gid, gname string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.ChatRoomEditNameReq{
		Gid:   gid,
		Gname: gname,
	}

	resultBody, err := client.DoPost("/ChatRoomEditName", req)
	if err != nil {
		return err
	}

	commonResult := model.ChatRoomEditNameResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ChatRoomEditName != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// ChatRoomEditAnnouncement 修改群聊公告（群主身份才可以使用） https://www.showdoc.com.cn/WeChatProject/8994689342748667
func ChatRoomEditAnnouncement(url string, req model.ChatRoomEditAnnouncementReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/ChatRoomEditAnnouncement", req)
	if err != nil {
		return err
	}

	commonResult := model.ChatRoomEditAnnouncementResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ChatRoomEditAnnouncement != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// ChatRoomVoip 群聊多人语音 https://www.showdoc.com.cn/WeChatProject/9019559882256250
func ChatRoomVoip(url string, gid string, wxidList []string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.ChatRoomVoipReq{
		Gid:      gid,
		WxidList: strings.Join(wxidList, ","),
	}

	resultBody, err := client.DoPost("/ChatRoomVoip", req)
	if err != nil {
		return err
	}

	commonResult := model.ChatRoomVoipResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ChatRoomVoip != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// QueryChatRoomOwnerWxid 查询群主微信ID https://www.showdoc.com.cn/WeChatProject/9019577323953766
func QueryChatRoomOwnerWxid(url string, gid string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return "", err
	}

	req := model.QueryChatRoomOwnerWxidReq{
		Gid: gid,
	}

	resultBody, err := client.DoPost("/QueryChatRoomOwnerWxid", req)
	if err != nil {
		return "", err
	}

	commonResult := model.QueryChatRoomOwnerWxidResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.OwnerWxid, nil
}

// GetChatroomMemberDetailInfo 网络获取群成员详细信息,可以获取群成员微信号 https://www.showdoc.com.cn/WeChatProject/9158592965427244
func GetChatroomMemberDetailInfo(url string, gid, wxid string) (*model.GetChatroomMemberDetailInfoResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return nil, err
	}

	req := model.GetChatroomMemberDetailInfoReq{
		Gid:  gid,
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/GetChatroomMemberDetailInfo", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.GetChatroomMemberDetailInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}

// GetChatrooMmemberDetail 批量获取群成员详细邀请信息,可以获取群成员是由谁邀请进来的 https://www.showdoc.com.cn/WeChatProject/9745666070753468
func GetChatrooMmemberDetail(url string, gid string) (*model.GetChatrooMmemberDetailResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return nil, err
	}

	req := model.GetChatrooMmemberDetailReq{
		Gid: gid,
	}

	resultBody, err := client.DoPost("/GetChatrooMmemberDetail", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.GetChatrooMmemberDetailResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}

// QueryChatRoomMemberCount 本地查询群成员总数(如果查询不到则需要先点一下群聊或者调用网络获取好友或群详细信息) https://www.showdoc.com.cn/WeChatProject/9019566491411017
func QueryChatRoomMemberCount(url string, gid string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return "", err
	}

	req := model.QueryChatRoomMemberCountReq{
		Gid: gid,
	}

	resultBody, err := client.DoPost("/QueryChatRoomMemberCount", req)
	if err != nil {
		return "", err
	}

	commonResult := model.QueryChatRoomMemberCountResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.Count, nil
}

// BatchGetChatRoomMemberWxid 批量获取群成员微信ID https://www.showdoc.com.cn/WeChatProject/9021207508542836
func BatchGetChatRoomMemberWxid(url string, gid string) ([]string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return nil, err
	}

	req := model.BatchGetChatRoomMemberWxidReq{
		Gid: gid,
	}

	resultBody, err := client.DoPost("/BatchGetChatRoomMemberWxid", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.BatchGetChatRoomMemberWxidResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	wxids := make([]string, 0)

	for _, data := range commonResult.Data {
		wxids = append(wxids, data.Wxid)
	}

	return wxids, nil
}

// QueryChatRoomMemberNickName 查询群成员昵称 https://www.showdoc.com.cn/WeChatProject/9019565404784998
func QueryChatRoomMemberNickName(url string, gid, wxid string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return "", err
	}

	req := model.QueryChatRoomMemberNickNameReq{
		Gid:  gid,
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/QueryChatRoomMemberNickName", req)
	if err != nil {
		return "", err
	}

	commonResult := model.QueryChatRoomMemberNickNameResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.Nickname, nil
}
