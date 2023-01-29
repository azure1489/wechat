package wechat

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/protobuf/ipc"
	"github.com/azure1489/wechat/util"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// TimelineGetFristPage 刷新并获取朋友圈第一页的内容，如果朋友圈有新动态则返回10条数据 https://www.showdoc.com.cn/WeChatProject/8929083282065703
func TimelineGetFristPage(url string) (*[]model.TimelineGetFristPageResultDataItem, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
	if err != nil {
		return nil, err
	}

	resultBody, err := client.DoPost("/TimelineGetFristPage", nil)
	if err != nil {
		return nil, err
	}

	// proto.Marshal(req)

	// fmt.Println("body:" + string(resultBody))

	// jsonpbMarshaler := &jsonpb.Marshaler{
	// 	EnumsAsInts:  true, // 是否将枚举值设定为整数，而不是字符串类型
	// 	EmitDefaults: true, // 是否将字段值为空的渲染到JSON结构中
	// 	OrigName:     true, // 是否使用原生的proto协议中的字段
	// }

	// protoMessage := proto.Message{}

	// jsonData, err := jsonpbMarshaler.MarshalToString(protoMessage)

	// 十六进制形式打印
	// for _, b := range resultBody {
	// 	// fmt.Println("%02X", b & 0xFF);
	// 	fmt.Printf("%#x\n", )

	// }

	// src = dst1
	maxDeLen := hex.DecodedLen(len(resultBody))
	dst := make([]byte, maxDeLen)
	// n, err := hex.Decode(dst1, src)

	// var dst []byte
	n, err := hex.Decode(dst, resultBody)
	if err != nil {
		return nil, err
	}

	// dst2 :=

	fmt.Printf("解码后的数据为:%s\n", string(dst[:n]))

	newData := &ipc.TimelineGetFristPageResult{}
	err = proto.UnmarshalOptions{
		DiscardUnknown: true,
		Merge:          true,
		AllowPartial:   true,
	}.Unmarshal(dst[:n], newData)
	if err != nil {
		// fmt.Println("unmarshal data err : ", err.Error())
		return nil, errors.Cause(err)
	}
	fmt.Println("unmarshal data : ", newData)

	// commonResult := model.TimelineGetFristPageResult{}
	// err = json.Unmarshal(resultBody, &commonResult)
	// if err != nil {
	// 	return nil, err
	// }

	// if commonResult.Havenewmsg != "1" {
	// 	return &[]model.TimelineGetFristPageResultDataItem{}, nil
	// }

	return nil, nil
}

// GetFriendTimeline 获取指定好友的首页朋友圈(返回最近10条记录) https://www.showdoc.com.cn/WeChatProject/9155297161590672
func GetFriendTimeline(url string, wxid string) (*[]model.GetFriendTimelineResultDataItem, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(url, timeout)
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
