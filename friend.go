package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
	"time"
)

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
