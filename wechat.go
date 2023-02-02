package wechat

type Wechat struct {
	Ip   string
	Port string
	Url  string
}

func NewWechat(ip, port, url string) *Wechat {
	return &Wechat{
		Url:  url,
		Ip:   ip,
		Port: port,
	}
}
