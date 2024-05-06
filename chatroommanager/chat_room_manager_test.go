package chatroommanager_test

import (
	"testing"
	"time"

	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/chatroommanager"
)

func TestQueryChatRoomMemberCount(t *testing.T) {

	config := wechat.WechatConfig{
		Ip:   "127.0.0.1",
		Port: "30001",
		// Url:  "https://wx.aworld.ltd/proxy",
		Url:           "https://proxy.aworld.ltd:9088/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",

		Timeout: time.Second * 60,
	}

	chatroomService := chatroommanager.NewChatRoomManagerService(&config)

	gid := "45117187985@chatroom"

	wxid, err := chatroomService.QueryChatRoomOwnerWxid(gid)
	if err != nil {
		t.Error(err)
	}

	t.Logf(" wxid:%s ", wxid)

	// wxid_svj371vacznb21
	count, err := chatroomService.QueryChatRoomMemberCount(gid)
	if err != nil {
		t.Error(err)
	}

	t.Logf(" count:%d ", count)
}

// func TestQueryChatRoomMemberCount(t *testing.T) {

// 	config := wechat.WechatConfig{
// 		Ip:   "127.0.0.1",
// 		Port: "30001",
// 		// Url:  "https://wx.aworld.ltd/proxy",
// 		Url:           "https://proxy.aworld.ltd:9088/proxy",
// 		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",

// 		Timeout: time.Second * 60,
// 	}

// 	chatroomService := chatroommanager.NewChatRoomManagerService(&config)

// 	gid := "45117187985@chatroom"

// 	wxid, err := chatroomService.QueryChatRoomOwnerWxid(gid)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Logf(" wxid:%s ", wxid)

// 	// wxid_svj371vacznb21
// 	count, err := chatroomService.GetChatrooMmemberDetail(gid)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Logf(" count:%d ", count)
// }
