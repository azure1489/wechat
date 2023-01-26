package wechat

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// AddNewFriend 添加新好友 https://www.showdoc.com.cn/WeChatProject/8961754683625021
func AddNewFriend(url string, req model.AddNewFriendReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/AddNewFriend", req)
	if err != nil {
		return err
	}

	commonResult := model.AddNewFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.AddNewFriend != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// DeleteFriend 删除好友 https://www.showdoc.com.cn/WeChatProject/8981501348113066
func DeleteFriend(url string, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.DeleteFriendReq{
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/DeleteFriend", req)
	if err != nil {
		return err
	}

	commonResult := model.DeleteFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.DeleteFriend != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// BlackFriend 拉黑好友 https://www.showdoc.com.cn/WeChatProject/9745604340432390
func BlackFriend(url string, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.BlackFriendReq{
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/BlackFriend", req)
	if err != nil {
		return err
	}

	commonResult := model.BlackFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.BlackFriend != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// VerifyFriend 同意好友申请 https://www.showdoc.com.cn/WeChatProject/8996001801759077
func VerifyFriend(url string, req model.VerifyFriendReq) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	resultBody, err := client.DoPost("/BlackFriend", req)
	if err != nil {
		return err
	}

	commonResult := model.VerifyFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.VerifyFriend != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// VoiceCallFriend 好友语音聊天 https://www.showdoc.com.cn/WeChatProject/9019318364312888
func VoiceCallFriend(url string, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.VoiceCallFriendReq{
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/VoiceCallFriend", req)
	if err != nil {
		return err
	}

	commonResult := model.VoiceCallFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.VoiceCallFriend != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// VideoCallFriend 好友视频聊天 https://www.showdoc.com.cn/WeChatProject/9019318842372957
func VideoCallFriend(url string, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.VideoCallFriendReq{
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/VideoCallFriend", req)
	if err != nil {
		return err
	}

	commonResult := model.VideoCallFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.VideoCallFriend != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// QueryBodyInfo 网络查询陌生人信息 https://www.showdoc.com.cn/WeChatProject/8982367120882654
func QueryBodyInfo(url string, serachwhat string) (*model.QueryBodyInfoResult, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return nil, err
	}

	req := model.QueryBodyInfoReq{
		SerachWhat: serachwhat,
	}

	resultBody, err := client.DoPost("/QueryBodyInfo", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.QueryBodyInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}

// EditFriendMark 修改好友备注 https://www.showdoc.com.cn/WeChatProject/9020497967635095
func EditFriendMark(url string, wxid, mark string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.EditFriendMarkReq{
		Wxid: wxid,
		Mark: mark,
	}

	resultBody, err := client.DoPost("/EditFriendMark", req)
	if err != nil {
		return err
	}

	commonResult := model.EditFriendMarkResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.EditFriendMark != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SetStarTag 标为星标朋友 https://www.showdoc.com.cn/WeChatProject/9078593012203121
func SetStarTag(url string, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SetStarTagReq{
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/SetStarTag", req)
	if err != nil {
		return err
	}

	commonResult := model.SetStarTagResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SetStarTag != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// UndoStarTag 取消星标朋友 https://www.showdoc.com.cn/WeChatProject/9078594588816258
func UndoStarTag(url string, wxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.UndoStarTagReq{
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/UndoStarTag", req)
	if err != nil {
		return err
	}

	commonResult := model.UndoStarTagResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.UndoStarTag != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SetContactRole 取消星标朋友 https://www.showdoc.com.cn/WeChatProject/9644721002325179
func SetContactRole(url string, wxid, role string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return err
	}

	req := model.SetContactRoleReq{
		Wxid: wxid,
		Role: role,
	}

	resultBody, err := client.DoPost("/SetContactRole", req)
	if err != nil {
		return err
	}

	commonResult := model.SetContactRoleResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SetContactRole != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
