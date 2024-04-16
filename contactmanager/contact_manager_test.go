package contactmanager_test

import (
	"testing"
	"time"

	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/contactmanager"
)

func TestGetFriendAndChatRoomList(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:   "127.0.0.1",
		Port: "30001",
		// Url:           "https://wx.aworld.ltd/proxy",
		Url:           "https://proxy.aworld.ltd:9088/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := contactmanager.NewContactManagerService(&config)
	// 0=返回所有 1=返回好友 2=返回群 3=返回公众号 4=企业微信联系人
	result, err := service.GetFriendAndChatRoomList("0")
	if err != nil {
		t.Error(err)
	}
	// t.Log(result)

	for _, v := range result.Chatroom {
		t.Log(v.Gname, v.Gid)
	}

	for _, v := range result.Friend {
		t.Log(v.Nickname, v.Wxid)
	}

}

func TestGetFriendOrChatroomDetailInfo(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "127.0.0.1",
		Port:          "30001",
		Url:           "https://wx.aworld.ltd/proxy",
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
		Ip:            "127.0.0.1",
		Port:          "30001",
		Url:           "https://wx.aworld.ltd/proxy",
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
