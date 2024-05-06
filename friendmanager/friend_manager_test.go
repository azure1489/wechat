package friendmanager_test

import (
	"testing"
	"time"

	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/friendmanager"
)

func TestAddNewFriend(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "127.0.0.1",
		Port:          "30001",
		Url:           "https://wx.aworld.ltd/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
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

func TestVerifyFriend(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:   "127.0.0.1",
		Port: "30001",
		// Url:  "https://wx.aworld.ltd/proxy",
		Url:           "https://proxy.aworld.ltd:9088/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",

		Timeout: time.Second * 60,
	}

	service := friendmanager.NewFriendManagerService(&config)

	req := friendmanager.VerifyFriendReq{
		V3:   "v3_020b3826fd03010000000000518200157d27fb000000501ea9a3dba12f95f6b60a0536a1adb6de1497253da9ecbc8f59982bf9031421f6850cc7a893603f0194942175abb482f8ce119ae093c0d71c50fbc6@stranger",
		V4:   "v4_000b708f0b04000001000000000099b0af7960c2a5c1494e6fba2c661000000050ded0b020927e3c97896a09d47e6e9ef3b5f81223de0269058f7bb10fd29c667dc2a857cd385e4897a06e4fb7660509f3007be809691974ad05a40da9320848b7d3a7758844d46fc6b88c16696cb40b1227bc9a1403817792b31a579a9296b97b26ef1c36bc90ecf4c3fd70f24c28077738cb004fbc7716@stranger",
		Role: "0",
		From: "17",
	}
	// wxid_svj371vacznb21
	err := service.VerifyFriend(&req)
	if err != nil {
		t.Error(err)
	}

}
