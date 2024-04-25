package loginmanager

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

// RefreshLoginQRCode 刷新登录二维码 https://www.showdoc.com.cn/WeChatProject/8966162223712985
func (l *LoginManagerServiceImpl) RefreshLoginQRCode() ([]byte, error) {

	resultBody, err := l.http.DoGet("/RefreshLoginQRCode")
	if err != nil {
		return nil, err
	}

	return resultBody, nil
}

// CreateRefreshLoginQRCode 创建登录二维码
func (l *LoginManagerServiceImpl) CreateRefreshLoginQRCode(jpegPath string) error {

	imgBytes, err := l.RefreshLoginQRCode()
	if err != nil {
		return err
	}

	// 创建一个新的缓冲区，并将字节切片写入缓冲区
	buf := bytes.NewBuffer(imgBytes)

	// 从缓冲区解码PNG图像
	img, err := png.Decode(buf)
	if err != nil {
		return err
	}

	// 创建一个新的图像，大小和原始图像一样，背景色为白色
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	draw.Draw(newImg, bounds, &image.Uniform{color.White}, image.Point{}, draw.Src)

	// 将原始图像绘制到新的图像上
	draw.Draw(newImg, bounds, img, image.Point{}, draw.Over)

	// 创建JPG文件
	jpgFile, err := os.Create(jpegPath)
	if err != nil {
		return err
	}
	defer jpgFile.Close()

	// 创建JPG编码选项，设置质量
	var opts jpeg.Options
	opts.Quality = 80

	// 将图像编码为JPG并写入文件
	err = jpeg.Encode(jpgFile, newImg, &opts)
	if err != nil {
		return err
	}

	return nil
}
