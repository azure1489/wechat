package wechat

import (
	"encoding/json"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// TimelineGetFristPage 刷新并获取朋友圈第一页的内容，如果朋友圈有新动态则返回10条数据 https://www.showdoc.com.cn/WeChatProject/8929083282065703
func (w *Wechat) TimelineGetFristPage() (*[]model.TimelineGetFristPageResultDataItem, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
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
