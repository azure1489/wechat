package loginmanager_test

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"testing"
	"time"

	// "github.com/Han-Ya-Jun/qrcode2console"
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/loginmanager"
)

// TestIsLoginStatus 测试获取微信登陆状态
func TestIsLoginStatus(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "127.0.0.1",
		Port:          "30001",
		Url:           "https://proxy.aworld.ltd:9088/proxy",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	result, err := service.IsLoginStatus()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

}

// TestGetWeChatProcessNumber 测试获取微信进程总数
func TestGetWeChatProcessNumber(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "172.16.153.221",
		Port:          "29998",
		Url:           "https://api.aworld.net.cn/wx",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	result, err := service.GetWeChatProcessNumber()
	if err != nil {
		t.Error(err)
	}

	for _, item := range result.List {
		t.Log("Par:", item.Par)
	}

	t.Log(result)

}

// GetPortOccupiedInfo 测试获取进程端口占用信息
func TestGetPortOccupiedInfo(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "172.16.153.221",
		Port:          "29998",
		Url:           "https://api.aworld.net.cn/wx",
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

	t.Log(result)
}

// TestStartWechat 测试启动更多微信
func TestStartWechat(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "172.16.153.169",
		Port:          "29999",
		Url:           "https://api.aworld.net.cn/wx",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	err := service.StartWechat(&loginmanager.StartWechatReq{
		StartPort: "30001",
		// WeChatPath: "C:\\soft\\Tencent\\WeChat\\WeChat.exe",
	})
	if err != nil {
		t.Error(err)
	}

}

// TestRefreshLoginQRCode 刷新并获取登录二维码
func TestRefreshLoginQRCode(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "172.16.153.221",
		Port:          "30002",
		Url:           "https://api.aworld.net.cn/wx",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	imgBytes, err := service.RefreshLoginQRCode()
	if err != nil {
		t.Error(err)
	}

	reader := bytes.NewReader(imgBytes)

	// 将字节数组转为图片
	img, _, err := image.Decode(reader)
	if err != nil {
		t.Fatal(err)
	}

	// 获取图片边界
	bounds := img.Bounds()

	// 创建一个带有白底的新图像
	whiteImage := image.NewRGBA(bounds)
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(whiteImage, whiteImage.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// 将PNG图片绘制到白底图像上
	draw.Draw(whiteImage, bounds, img, bounds.Min, draw.Over)

	// 字节转成图片，显示二维码图片
	// 2. 将字节缓冲区写入临时文件（因为GTK需要从文件中读取）
	outFile, err := os.CreateTemp("", "example.*.jpg")
	if err != nil {
		println("Could not create temp file:", err)
		return
	}
	defer outFile.Close()

	// 将图像编码为JPG格式
	var opts jpeg.Options
	opts.Quality = 80
	err = jpeg.Encode(outFile, whiteImage, &opts)
	if err != nil {
		panic(err)
	}

	fmt.Println(outFile.Name())

	// qr := qrcode2console.NewQRCode2ConsoleWithPath(outFile.Name())

	// qr.Output()
}

// TestGetSelfLoginInfo 测试获取个人详细信息
func TestGetSelfLoginInfo(t *testing.T) {
	config := wechat.WechatConfig{
		Ip:            "172.16.153.221",
		Port:          "30001",
		Url:           "https://api.aworld.net.cn/wx",
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
		Ip:            "172.16.153.221",
		Port:          "30002",
		Url:           "https://api.aworld.net.cn/wx",
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
		Ip:            "172.16.153.221",
		Port:          "30002",
		Url:           "https://api.aworld.net.cn/wx",
		PublicKeyPath: "/Users/azure/git/go-project/text-to-silk/proxy-public.pem",
		Timeout:       time.Second * 60,
	}

	service := loginmanager.NewLoginManagerService(&config)
	err := service.TerminateThisWeChat()
	if err != nil {
		t.Error(err)
	}
}
