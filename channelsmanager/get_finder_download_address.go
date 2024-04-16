package channelsmanager

import "encoding/json"

// 返回参数说明

// 参数名	类型	说明
// DownloadAddress	string	视频的下载URL地址
// thumb_url	string	视频预览图
// title	string	视频标题
// md5	string	视频MD5值
// media_type	string	视频类别
// filesize	string	视频大小
// play_len	string	视频播放时长(单位:秒)
// width	string	视频宽
// height	string	视频高
//
//	{
//	    "DownloadAddress": "https://finder.video.qq.com/251/20302/stodownload?encfilekey=Cvvj5Ix3eewK0tHtibORqcsqchXNh0Gf3sJcaYqC2rQDk90hTVCJVGmEeRKe6yadMxbm2x0CUVRsAo0UC1qjCV3mU2IGtsX4H67iagMU1CHIXWN1BLXwCaJa61D1F0icomib&token=AxricY7RBHdXo2Dic6YY8diblOmGYgRdCIjqjariavcxib9XXCRfBia3Jx7fQL4ssgdQh8ic3mUZCvd0Eg&idx=1&adaptivelytrans=0&bizid=1023&dotrans=3071&hy=SH&m=&finderextscene=6",
//	    "thumb_url": "https://finder.video.qq.com/251/20304/stodownload?encfilekey=rjD5jyTuFrIpZ2ibE8T7YmwgiahniaXswqzCLxNF3EyRTmSgNzOpu51LNv9PCrTaVvXtt3Eibc5oYJiapSTuRxeTzibxJypPP4Jialou3ibuSo1z32r8k7Ua0ia8kYw&token=6xykWLEnztKIzBicPuvgFxuXnXDmAPIT66WBEPHoBw5Iy3uXspEMborrAozrphajia&idx=1&adaptivelytrans=0&bizid=1023&dotrans=0&hy=SH&m=&scene=0&finderextscene=6",
//	    "title": "你脑子有问题吧，这么仔细干嘛！",
//	    "md5": "7e676e78cae94b216463d4b77dce98f7",
//	    "media_type": "4",
//	    "filesize": "48697949",
//	    "play_len": "38",
//	    "width": "1080",
//	    "height": "1920"
//	}
type GetFinderDownloadAddressRes struct {
	DownloadAddress string `json:"DownloadAddress"`
	ThumbUrl        string `json:"thumb_url"`
	Title           string `json:"title"`
	Md5             string `json:"md5"`
	MediaType       string `json:"media_type"`
	Filesize        string `json:"filesize"`
	PlayLen         string `json:"play_len"`
	Width           string `json:"width"`
	Height          string `json:"height"`
}

// 参数名	类型	说明
// objectId	string	视频objectId
// objectNonceId	string	视频objectNonceId
// 获取视频下载地址
func (l *ChannelsManagerServiceImpl) GetFinderDownloadAddress(objectId string, objectNonceId string) (*GetFinderDownloadAddressRes, error) {

	req := map[string]interface{}{
		"objectId":      objectId,
		"objectNonceId": objectNonceId,
	}

	resultBody, err := l.http.DoPost("/GetFinderDownloadAddress", req)
	if err != nil {
		return nil, err
	}

	commonResult := GetFinderDownloadAddressRes{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
