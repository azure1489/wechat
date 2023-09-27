package loginmanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type LoginManagerService interface {

	// 获取微信进程总数 https://www.showdoc.com.cn/WeChatProject/9794632555004152
	GetWeChatProcessNumber() (*GetWeChatProcessNumberResult, error)

	// Get_PortOccupiedInfo 获取进程端口占用信息 https://www.showdoc.com.cn/WeChatProject/10128598529687136
	GetPortOccupiedInfo(*PortOccupiedInfoReq) (string, error)

	// AgainStartWeChat 启动更多微信 https://www.showdoc.com.cn/WeChatProject/9063540299207309
	StartWechat(*StartWechatReq) error

	// ClickLoginButton 点击登陆微信 https://www.showdoc.com.cn/WeChatProject/10373442896174744
	ClickLoginButton() error

	// RefreshLoginQRCode 刷新登录二维码 https://www.showdoc.com.cn/WeChatProject/8966162223712985
	RefreshLoginQRCode() ([]byte, error)

	// GetLoginQRCode 获取登录二维码图片 https://www.showdoc.com.cn/WeChatProject/9026367939952187
	GetLoginQRCode() (string, error)

	// IsLoginStatus 获取微信登陆状态 https://www.showdoc.com.cn/WeChatProject/8991915897471323
	IsLoginStatus() (*IsLoginStatusResult, error)

	// GetSelfLoginInfo 获取个人详细信息 https://www.showdoc.com.cn/WeChatProject/8929111706614173
	GetSelfLoginInfo() (*GetSelfLoginInfoResult, error)

	// Logout 退出微信 https://www.showdoc.com.cn/WeChatProject/9008665345790557
	Logout() error

	// TerminateThisWeChat 结束微信 https://www.showdoc.com.cn/WeChatProject/9214210657048561
	TerminateThisWeChat() error
}

type LoginManagerServiceImpl struct {
	// config *wechat.WechatConfig
	http common.HttpClientService
}

func NewLoginManagerService(config *wechat.WechatConfig) LoginManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.Secret, config.Timeout)

	return &LoginManagerServiceImpl{
		// config: config,
		http: httpClientService,
	}
}
