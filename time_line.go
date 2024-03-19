package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// TimelineGetFristPage 刷新并获取朋友圈第一页的内容，如果朋友圈有新动态则返回10条数据 https://www.showdoc.com.cn/WeChatProject/8929083282065703
func (w *Wechat) TimelineGetFristPage() (*[]model.TimelineGetFristPageResultDataItem, error) {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return nil, err
	}

	resultBody, err := client.DoPost("/TimelineGetFristPage", nil)
	if err != nil {
		return nil, err
	}

	commonResult := model.TimelineGetFristPageResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	if commonResult.Count == "0" {
		return &[]model.TimelineGetFristPageResultDataItem{}, nil
	}

	return &commonResult.Data, nil
}

// TimelineGetNextPage 朋友圈获取下一页内容 https://www.showdoc.com.cn/WeChatProject/8929087905557074
func (w *Wechat) TimelineGetNextPage(id string) (*[]model.TimelineGetFristPageResultDataItem, error) {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return nil, err
	}

	req := model.TimelineGetNextPageReq{
		Id: id,
	}

	resultBody, err := client.DoPost("/TimelineGetNextPage", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.TimelineGetFristPageResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	if commonResult.Count == "0" {
		return &[]model.TimelineGetFristPageResultDataItem{}, nil
	}

	return &commonResult.Data, nil
}

// GetFriendTimeline 获取指定好友的首页朋友圈(返回最近10条记录) https://www.showdoc.com.cn/WeChatProject/9155297161590672
func (w *Wechat) GetFriendTimeline(wxid string) (*[]model.GetFriendTimelineResultDataItem, error) {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return nil, err
	}

	req := model.GetFriendTimelineReq{
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/GetFriendTimeline", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.GetFriendTimelineResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult.Data, nil
}

// SendTimeline 转发朋友圈 https://www.showdoc.com.cn/WeChatProject/9137221496703083
func (w *Wechat) SendTimeline(sendPyqXml string) error {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return err
	}

	req := model.SendTimelineReq{
		SendPyqXml: sendPyqXml,
	}

	resultBody, err := client.DoPost("/SendTimeline", req)
	if err != nil {
		return err
	}

	commonResult := model.SendTimelineResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendTimeline != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// TimelineLike 朋友圈点赞 https://www.showdoc.com.cn/WeChatProject/8929089716176919
func (w *Wechat) TimelineLike(id string) error {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return err
	}

	req := model.TimelineLikeReq{
		ID: id,
	}

	resultBody, err := client.DoPost("/TimelineLike", req)
	if err != nil {
		return err
	}

	commonResult := model.TimelineLikeResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.TimelineLike != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// TimelineUndoLike 朋友圈取消点赞 https://www.showdoc.com.cn/WeChatProject/8929091706495732
func (w *Wechat) TimelineUndoLike(id string) error {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return err
	}

	req := model.TimelineUndoLikeReq{
		ID: id,
	}

	resultBody, err := client.DoPost("/TimelineUndoLike", req)
	if err != nil {
		return err
	}

	commonResult := model.TimelineUndoLikeResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.TimelineUndoLike != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// TimelineComment 朋友圈发表评论内容 https://www.showdoc.com.cn/WeChatProject/8929091996936675
func (w *Wechat) TimelineComment(id, msg string) error {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return err
	}

	req := model.TimelineCommentReq{
		ID:  id,
		Msg: msg,
	}

	resultBody, err := client.DoPost("/TimelineComment", req)
	if err != nil {
		return err
	}

	commonResult := model.TimelineCommentResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.TimelineComment != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// TimelineDeleteComment 删除朋友圈评论内容 https://www.showdoc.com.cn/WeChatProject/8929092413429617
func (w *Wechat) TimelineDeleteComment(id, index string) error {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return err
	}

	req := model.TimelineDeleteCommentReq{
		ID:    id,
		Index: index,
	}

	resultBody, err := client.DoPost("/TimelineDeleteComment", req)
	if err != nil {
		return err
	}

	commonResult := model.TimelineDeleteCommentResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.TimelineDeleteComment != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// SwitchTimelineComment 开关朋友圈评论功能 https://www.showdoc.com.cn/WeChatProject/9020961696935450
func (w *Wechat) SwitchTimelineComment(option string) error {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return err
	}

	req := model.SwitchTimelineCommentReq{
		Option: option,
	}

	resultBody, err := client.DoPost("/SwitchTimelineComment", req)
	if err != nil {
		return err
	}

	commonResult := model.SwitchTimelineCommentResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SwitchTimelineComment != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}

// TimelineUploadPic 朋友圈上传图片 https://www.showdoc.com.cn/WeChatProject/9851577390989318
func (w *Wechat) TimelineUploadPic(option string) (*model.TimelineUploadPicResult, error) {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return nil, err
	}

	req := model.SwitchTimelineCommentReq{
		Option: option,
	}

	resultBody, err := client.DoPost("/TimelineUploadPic", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.TimelineUploadPicResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
