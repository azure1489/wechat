package sendmsgmanager_test

import (
	"testing"
	"time"

	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/messagemanager/sendmsgmanager"
)

func TestSendTextMsg(t *testing.T) {
	config := wechat.WechatConfig{
		// Ip:            "127.0.0.1",
		// Port:          "30001",
		// Url:           "https://proxy.aworld.ltd:9088/proxy",
		Ip:            "127.0.0.1",
		Port:          "30001",
		Url:           "https://wx.aworld.ltd/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := sendmsgmanager.NewSendMsgManagerService(&config)
	err := service.SendTextMsg("ccs-lemon", "这不是歧视其他省吗？")
	if err != nil {
		t.Error(err)
	}

}
