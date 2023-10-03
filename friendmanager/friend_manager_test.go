package friendmanager_test

import (
	"testing"
	"time"

	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/friendmanager"
)

func TestAddNewFriend(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:      "172.16.153.169",
		Port:    "30001",
		Url:     "https://api.aworld.net.cn/wx",
		Secret:  "81fb6a51e232411c09575bb96bf71675980da0ac",
		Timeout: time.Second * 60,
	}

	service := friendmanager.NewFriendManagerService(&config)

	req := friendmanager.AddNewFriendReq{
		V3WxId: "wxid_svj371vacznb21",
		Desc:   "你好",
		Type:   "30",
		Role:   "0",
	}
	// wxid_svj371vacznb21
	commonResult, err := service.AddNewFriend(req)
	if err != nil {
		t.Error(err)
	}

	t.Log(commonResult)
}