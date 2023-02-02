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
func (w *Wechat) CreateChatRoom(wxidList []string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) AddFriendToChatRoom(gid string, wxidList []string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) InviteFriendToChatRoom(gid, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) AgreeJoinGroup(groupAccessUrl string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) ChatRoomEditMark(gid, mark string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) ChatRoomEditName(gid, gname string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) ChatRoomEditAnnouncement(req model.ChatRoomEditAnnouncementReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) ChatRoomVoip(gid string, wxidList []string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) QueryChatRoomOwnerWxid(gid string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) GetChatroomMemberDetailInfo(gid, wxid string) (*model.GetChatroomMemberDetailInfoResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) GetChatrooMmemberDetail(gid string) (*model.GetChatrooMmemberDetailResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) QueryChatRoomMemberCount(gid string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) BatchGetChatRoomMemberWxid(gid string) ([]string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
func (w *Wechat) QueryChatRoomMemberNickName(gid, wxid string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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

// ChatRoomMemberBatchDelete 群成员批量删除 https://www.showdoc.com.cn/WeChatProject/8993612184500147
func (w *Wechat) ChatRoomMemberBatchDelete(gid string, wxidList []string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.ChatRoomMemberBatchDeleteReq{
		Gid:      gid,
		WxidList: strings.Join(wxidList, ","),
	}

	resultBody, err := client.DoPost("/ChatRoomMemberBatchDelete", req)
	if err != nil {
		return err
	}

	commonResult := model.ChatRoomMemberBatchDeleteResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ChatRoomMemberBatchDelete != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// ChangeChatRoomNicknameBySelf 修改自己的群昵称 https://www.showdoc.com.cn/WeChatProject/9019580625786730
func (w *Wechat) ChangeChatRoomNicknameBySelf(req model.ChangeChatRoomNicknameBySelfReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/ChangeChatRoomNicknameBySelf", req)
	if err != nil {
		return err
	}

	commonResult := model.ChangeChatRoomNicknameBySelfResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ChangeChatRoomNicknameBySelf != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// QuitChatRoom 退出群聊 https://www.showdoc.com.cn/WeChatProject/9002651226705549
func (w *Wechat) QuitChatRoom(gid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.QuitChatRoomReq{
		Gid: gid,
	}

	resultBody, err := client.DoPost("/QuitChatRoom", req)
	if err != nil {
		return err
	}

	commonResult := model.QuitChatRoomResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.QuitChatRoom != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// DisplayChatRoomMemberNickName 显示群成员昵称 https://www.showdoc.com.cn/WeChatProject/9063349491429947
func (w *Wechat) DisplayChatRoomMemberNickName(gid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.DisplayChatRoomMemberNickNameReq{
		Gid: gid,
	}

	resultBody, err := client.DoPost("/DisplayChatRoomMemberNickName", req)
	if err != nil {
		return err
	}

	commonResult := model.DisplayChatRoomMemberNickNameResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.DisplayChatRoomMemberNickName != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// TurnOffDisplayChatRoomMemberNickName 关闭显示群成员昵称 https://www.showdoc.com.cn/WeChatProject/9063350057038457
func (w *Wechat) TurnOffDisplayChatRoomMemberNickName(gid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.TurnOffDisplayChatRoomMemberNickNameReq{
		Gid: gid,
	}

	resultBody, err := client.DoPost("/TurnOffDisplayChatRoomMemberNickName", req)
	if err != nil {
		return err
	}

	commonResult := model.TurnOffDisplayChatRoomMemberNickNameResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.TurnOffDisplayChatRoomMemberNickName != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// CollapseChat 折叠该群聊 https://www.showdoc.com.cn/WeChatProject/9063366035200767
func (w *Wechat) CollapseChat(gid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.CollapseChatReq{
		Gid: gid,
	}

	resultBody, err := client.DoPost("/CollapseChat", req)
	if err != nil {
		return err
	}

	commonResult := model.CollapseChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.CollapseChat != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// UncollapseChat 取消折叠该群聊 https://www.showdoc.com.cn/WeChatProject/9063366996691400
func (w *Wechat) UncollapseChat(gid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return err
	}

	req := model.UncollapseChatReq{
		Gid: gid,
	}

	resultBody, err := client.DoPost("/UncollapseChat", req)
	if err != nil {
		return err
	}

	commonResult := model.UncollapseChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.UncollapseChat != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
