package contactmanager_test

import (
	"testing"
	"time"

	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/contactmanager"
)

func TestGetFriendAndChatRoomList(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:      "172.16.153.221",
		Port:    "30001",
		Url:     "https://api.aworld.net.cn/wx",
		Secret:  "81fb6a51e232411c09575bb96bf71675980da0ac",
		Timeout: time.Second * 60,
	}

	service := contactmanager.NewContactManagerService(&config)
	result, err := service.GetFriendAndChatRoomList("0")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
