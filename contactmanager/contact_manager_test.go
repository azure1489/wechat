package contactmanager_test

import (
	"testing"
	"time"

	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/contactmanager"
)

func TestGetFriendAndChatRoomList(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "172.16.153.169",
		Port:          "30001",
		Url:           "https://api.aworld.net.cn/wx",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := contactmanager.NewContactManagerService(&config)
	result, err := service.GetFriendAndChatRoomList("0")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestGetFriendOrChatroomDetailInfo(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "172.16.153.221",
		Port:          "30001",
		Url:           "https://api.aworld.net.cn/wx",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := contactmanager.NewContactManagerService(&config)
	// wxid_svj371vacznb21 wxid_yaalopa9vcka21
	result, err := service.GetFriendOrChatroomDetailInfo("wxid_50l6xqg1bmaq22")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestBeforeTransfer(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "172.16.153.169",
		Port:          "30001",
		Url:           "https://api.aworld.net.cn/wx",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := contactmanager.NewContactManagerService(&config)
	// wxid_svj371vacznb21 wxid_yaalopa9vcka21
	result, err := service.BeforeTransfer("wxid_3435654360314")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
