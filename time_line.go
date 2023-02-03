package wechat

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// TimelineGetFristPage 刷新并获取朋友圈第一页的内容，如果朋友圈有新动态则返回10条数据 https://www.showdoc.com.cn/WeChatProject/8929083282065703
func (w *Wechat) TimelineGetFristPage(url string) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return "", err
	}

	resultBody, err := client.DoPost("/TimelineGetFristPage", nil)
	if err != nil {
		return "", err
	}

	// proto.Unmarshal()

	// m := mahonia.NewDecoder("utf8")

	// m.Translate(resultBody,true)

	maxDeLen := hex.DecodedLen(len(resultBody))
	dst := make([]byte, maxDeLen)
	n, err := hex.Decode(dst, resultBody)
	if err != nil {
		return "", err
	}
	decodeData := dst[:n]

	// bArr, err := hex.DecodeString(string(resultBody))
	// if err != nil {
	// 	return nil, err
	// }

	// e := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	// es, _, err := transform.Bytes(e.NewEncoder(), decodeData)
	// if err != nil {
	// 	return nil, err
	// }

	// enc := mahonia.NewEncoder("gbk")
	//converts a  string from UTF-8 to gbk encoding.
	// fmt.Println(enc.ConvertString(string(decodeData)))

	// fmt.Println(string(decodeData))

	// r := bytes.Runes(bArr)

	// // res := ""
	// for _, b := range r {
	// 	// res = fmt.Sprintf("%s%.8b", res, b)
	// 	fmt.Print(string(b))
	// }

	// srcCoder := mahonia.NewDecoder(text)
	// tagCoder := mahonia.NewDecoder(tagCode)
	// _, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	// result := string(cdata)

	// text = ConvertToString(text, "utf-8", "utf-8")

	// fmt.Println("body:" + text)

	// proto.Marshal(req)

	// fmt.Println("body:\n" + text)

	// jsonpbMarshaler := &jsonpb.Marshaler{
	// 	EnumsAsInts:  true, // 是否将枚举值设定为整数，而不是字符串类型
	// 	EmitDefaults: true, // 是否将字段值为空的渲染到JSON结构中
	// 	OrigName:     true, // 是否使用原生的proto协议中的字段
	// }

	// protoMessage := proto.Message{}

	// jsonData, err := jsonpbMarshaler.MarshalToString(protoMessage)

	// src = dst1
	// maxDeLen := hex.DecodedLen(len(resultBody))
	// dst := make([]byte, maxDeLen)
	// // n, err := hex.Decode(dst1, src)

	// // var dst []byte
	// n, err := hex.Decode(dst, resultBody)
	// if err != nil {
	// 	return nil, err
	// }

	// dst2 :=

	fmt.Println("解码后的数据为:")

	fmt.Println(" ---------- ")
	// 十六进制形式打印
	// for _, b := range dst[:n] {
	// 	fmt.Printf("uint8 UTF-8 %d \n", b)
	// 	fmt.Printf("utf-8 %q \n", b)
	// }

	// utf8.DecodeRune()

	// utf8bs := make([]byte, 0)
	// bs := make([]byte, 4)
	// for index, b := range dst[:n] {
	// 	if (index + 1) % 4 == 0 {

	// 	} else {

	// 	}
	// }

	// testInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

	// testInt2 := make([][]byte, 0)
	// testInt3 := make([]byte, 0)
	// byteLan := 4

	// for index, r := range dst[:n] {
	// 	// fmt.Println("(index+1)%4:", (index+1)%4)
	// 	if (index+1)%byteLan != 0 {
	// 		testInt3 = append(testInt3, r)
	// 	} else {
	// 		testInt3 = append(testInt3, r)
	// 		testInt2 = append(testInt2, testInt3)
	// 		testInt3 = make([]byte, 0)
	// 	}
	// 	if index == len(dst[:n])-1 {
	// 		testInt2 = append(testInt2, testInt3)
	// 	}
	// }

	// for _, r := range testInt2 {
	// 	fmt.Print(string(r))
	// }

	// s := dst[:n]

	// t := make([]rune, utf16.(s))
	// i := 0
	// for len(s) > 0 {
	// 	r, l := utf8.DecodeRune(s)
	// 	t[i] = r
	// 	i++
	// 	s = s[l:]
	// }

	// textRunes := t

	// for _, test := range textRunes {
	// 	fmt.Print(string(test))
	// }

	// fmt.Print(string(t))

	// utf8bs := make([]byte, 0)
	// for _, r := range textRune {
	// 	bs := make([]byte, 4)
	// 	w := utf8.EncodeRune(bs, r)
	// 	utf8bs = append(utf8bs, bs[:w]...)
	// }
	// fmt.Println(string(utf8bs))

	// test,size := utf8.DecodeRune(dst[:n])
	// r, w := utf8.DecodeRune(dst)
	// fmt.Printf("%x ", r)
	// fmt.Print(string(r))
	// //     // fmt.Printf("%x ", r)
	// // 	fmt.Print(string(v))

	// 	decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	// 	// 	bs2, err := decoder.Bytes(bs[:])
	// for _, b := range dst[:n] {
	// // 	// fmt.Printf("%d:%s|", k, string(v))
	// // 	fmt.Println(utf8.DecodeLastRune(b))
	// }

	// text := string(dst[:n])

	// text = ConvertToString(text, "utf-8", "utf-16")

	// fmt.Println(text)

	// textRune := []rune(text)
	// // textLen := len(textRune)
	// // fmt.Println("len:" + strconv.FormatInt(int64(textLen), 10))
	// // for _, v := range textRune {

	// // 	fmt.Print(string(v))
	// // }

	// // bs := []byte(text)
	// // for len(bs) > 0 {
	// //     r, w := utf8.DecodeRune(bs)
	// //     // fmt.Printf("%x ", r)
	// // 	fmt.Print(string(v))
	// //     bs = bs[w:]
	// // }

	// utf8bs := make([]byte, 0)
	// for _, r := range textRune {
	// 	bs := make([]byte, 4)
	// 	w := utf8.EncodeRune(bs, r)
	// 	utf8bs = append(utf8bs, bs[:w]...)
	// }
	// fmt.Println(string(utf8bs))

	fmt.Println("\n ---------- ")

	// dataInfo := &ipc.TimelineGetFristPageResult{}
	// err = proto.UnmarshalOptions{
	// 	Merge:          true,
	// 	AllowPartial:   true,
	// 	DiscardUnknown: true,
	// }.Unmarshal(decodeData, dataInfo)
	// if err != nil {
	// 	// fmt.Println("unmarshal data err : ", err.Error())
	// 	return nil, errors.Cause(err)
	// }
	// fmt.Println("unmarshal data : ", dataInfo)

	// commonResult := model.TimelineGetFristPageResult{}
	// err = json.Unmarshal(resultBody, &commonResult)
	// if err != nil {
	// 	return nil, err
	// }

	// if commonResult.Havenewmsg != "1" {
	// 	return &[]model.TimelineGetFristPageResultDataItem{}, nil
	// }

	return string(decodeData), nil
}

// GetFriendTimeline 获取指定好友的首页朋友圈(返回最近10条记录) https://www.showdoc.com.cn/WeChatProject/9155297161590672
func (w *Wechat) GetFriendTimeline(wxid string) (*[]model.GetFriendTimelineResultDataItem, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return nil, err
	}

	req := model.GetFriendTimelineReq{
		Wxid: wxid,
	}

	resultBody, err := client.DoPost("/GetFriendTimeline", req)
	if err != nil {
		return nil, err
	}

	commonResult := model.GetFriendTimelineResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult.Data, nil
}
