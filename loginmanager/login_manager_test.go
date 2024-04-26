package loginmanager_test

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"
	"time"

	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/loginmanager"
)

// TestIsLoginStatus 测试获取微信登陆状态
func TestIsLoginStatus(t *testing.T) {

	config2 := wechat.WechatConfig{
		Ip:   "127.0.0.1",
		Port: "29998",
		// Url:  "https://proxy.aworld.ltd:9088/proxy",
		Url:           "https://wx.aworld.ltd/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service2 := loginmanager.NewLoginManagerService(&config2)
	result2, err := service2.GetWeChatPort()
	if err != nil {
		t.Error(err)
	}

	for _, port := range result2 {
		t.Log("port:", port)
		config := wechat.WechatConfig{
			Ip:   "127.0.0.1",
			Port: port,
			// Url:           "https://proxy.aworld.ltd:9088/proxy",
			Url:           "https://wx.aworld.ltd/proxy",
			PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
			Timeout:       time.Second * 60,
		}

		service := loginmanager.NewLoginManagerService(&config)
		result, err := service.IsLoginStatus()
		if err != nil {
			t.Error(err)
		}
		// 当前微信在线状态 OnlineStatus: 0=请扫码登陆 1=请在手机上完成登录 2=正在登陆中 3=完登陆完成 4=正在退出微信 5=点击进入微信
		t.Log("OnlineStatus:", result.OnlineStatus, ", LoginLoading:", result.LoginLoading, ", SelfWxid:", result.SelfWxid, ", NickName:", result.NickName)
	}

}

// TestGetWeChatProcessNumber 测试获取微信进程总数
func TestGetWeChatProcessNumber(t *testing.T) {

	config := wechat.WechatConfig{
		Ip:   "127.0.0.1",
		Port: "29998",
		// Url:  "https://proxy.aworld.ltd:9088/proxy",
		Url:           "https://wx.aworld.ltd/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	result, err := service.GetWeChatProcessNumber()
	if err != nil {
		t.Error(err)
	}

	log.Println("1总进程数：", result.TotalNum)

	// pids := make([]string, 0)
	for _, item := range result.List {
		t.Log("Par:", item.Par)
		t.Log("PID:", item.PID)
		t.Log("ProcessName:", item.ProcessName)
		t.Log("Port:", item.Port)
		// pids = append(pids, strconv.Itoa(item.PID))
	}

	// urlStr := "https://proxy.aworld.ltd:9088/process-ports"

	// u, err := url.Parse(urlStr)
	// if err != nil {
	// 	panic(err)
	// }
	// t.Log("hostname:" + u.Hostname())
	// // url := "https://wx.aworld.ltd/process-ports"
	// publicKeyPath := "/Users/azure/git/go-project/test-project/proxy-public.pem"

	// ports, err := service.GetProcessPorts(urlStr, time.Second*60, pids, publicKeyPath)
	// if err != nil {
	// 	t.Error(err)
	// }

	// for _, port := range ports {
	// 	t.Log("port:", port)
	// }

	// t.Log(result)

}

// GetPortOccupiedInfo 测试获取进程端口占用信息
func TestGetPortOccupiedInfo(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:   "127.0.0.1",
		Port: "29998",
		// Url:           "https://proxy.aworld.ltd:9088/proxy",
		Url:           "https://wx.aworld.ltd/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	result, err := service.GetPortOccupiedInfo(&loginmanager.PortOccupiedInfoReq{
		CheckPort: "30001",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log("result:" + result)
}

// TestStartWechat 测试启动更多微信
func TestStartWechat(t *testing.T) {

	config := wechat.WechatConfig{
		Ip:   "127.0.0.1",
		Port: "29998",
		// Url:           "https://wx.aworld.ltd/proxy",
		Url:           "https://proxy.aworld.ltd:9088/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	err := service.StartWechat(&loginmanager.StartWechatReq{
		StartPort:  "30003",
		WeChatPath: "C:\\soft\\Tencent\\WeChat\\",
	})
	if err != nil {
		t.Error(err)
	}

}

// TestRefreshLoginQRCode 刷新并获取登录二维码
func TestRefreshLoginQRCode(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:   "127.0.0.1",
		Port: "30001",
		// Url:  "https://wx.aworld.ltd/proxy",
		Url:           "https://proxy.aworld.ltd:9088/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	imgBytes, err := service.RefreshLoginQRCode()
	if err != nil {
		t.Error(err)
	}

	// 创建一个新的缓冲区，并将字节切片写入缓冲区
	buf := bytes.NewBuffer(imgBytes)

	// 从缓冲区解码PNG图像
	img, err := png.Decode(buf)
	if err != nil {
		panic(err)
	}
	// 创建一个新的图像，大小和原始图像一样，背景色为白色
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	draw.Draw(newImg, bounds, &image.Uniform{color.White}, image.Point{}, draw.Src)

	// 将原始图像绘制到新的图像上
	draw.Draw(newImg, bounds, img, image.Point{}, draw.Over)

	// 创建JPG文件
	jpgFile, err := os.Create("/Users/azure/Downloads/output10.jpg")
	if err != nil {
		panic(err)
	}
	defer jpgFile.Close()

	// 创建JPG编码选项，设置质量
	var opts jpeg.Options
	opts.Quality = 80

	// 将图像编码为JPG并写入文件
	err = jpeg.Encode(jpgFile, newImg, &opts)
	if err != nil {
		panic(err)
	}
}

// TestGetSelfLoginInfo 测试获取个人详细信息
func TestGetSelfLoginInfo(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "172.16.153.221",
		Port:          "30001",
		Url:           "https://wx.aworld.ltd/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	result, err := service.GetSelfLoginInfo()
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

// TestLogout 测试退出微信
func TestLogout(t *testing.T) {

	config := wechat.WechatConfig{
		Ip:            "127.0.0.1",
		Port:          "30003",
		Url:           "https://proxy.aworld.ltd:9088/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	err := service.Logout()
	if err != nil {
		t.Error(err)
	}
}

// TestTerminateThisWeChat 测试结束微信
func TestTerminateThisWeChat(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:   "127.0.0.1",
		Port: "30001",
		// Url:           "https://wx.aworld.ltd/proxy",
		Url:           "https://proxy.aworld.ltd:9088/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	err := service.TerminateThisWeChat()
	if err != nil {
		t.Error(err)
	}
}
