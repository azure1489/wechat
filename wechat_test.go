package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

// func GetMD5Encode(data string) string {
// 	h := md5.New()
// 	h.Write([]byte(data))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func TestGetFriendAndChatRoomList(t *testing.T) {

	//
	//{"app_id":"uEVq0SDj34HwVltpNKzdgxmK",
	//	"msg_id":"1910117248440097833984",
	//	"tid":"3708064405066220800","data":{"company_id":1635,"shop_id":1910,"branch_id":4,"tid":"3708064405066220800","oid":"3708064405066220800001","express_name":"%e9%a1%ba%e4%b8%b0%e9%80%9f%e8%bf%90","express_id":"7","waybill_no":"SF1651895583219","express_date":"12/13/2022 14:45:23","operate_user":"8642"},"timestamp":"1670942746960"}
	//

	url := "https://app.aworld.ltd:9112/api"
	ip := "127.0.0.1"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	friendAndChatRooms, err := NewWechat(ip, port, url, appSecret).GetFriendAndChatRoomList()
	if err != nil {
		t.Error(err)
		return
	}

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(friendAndChatRooms)

	fmt.Println("data:", bf.String())

	//if len(friendAndChatRooms.Friends) > 0 {
	//	for _, friend := range friendAndChatRooms.Friends {
	//		friendStatus, err := CheckFriendStatus(friend.Wxid)
	//		if err != nil {
	//			t.Error(err)
	//			return
	//		}
	//		if friendStatus != 1 {
	//			log.Println("wxid:", friend.Wxid)
	//			log.Println("非好友 friendStatus:", friendStatus)
	//			log.Println("昵称:", friend.Nickname)
	//			log.Println("头像:", friend.Headimg)
	//		}
	//	}
	//}
}

func TestSendTextMsg(t *testing.T) {
	// url := "https://mp.aworld.ltd/wx/api"
	// ip := "127.0.0.1"
	// port := "30001"
	// appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"

	url := "http://10.211.4.239:30001"
	err := NewWechatEncryption(url).SendTextMsg("filehelper", "e123")
	if err != nil {
		t.Error(err)
		return
	}
}

// func TestSendVoiceMsg(t *testing.T) {

// 	// Read Silk file
// 	silkFile := "/Users/yuanhua/Downloads/test02.silk"
// 	data, err := ioutil.ReadFile(silkFile)
// 	if err != nil {
// 		fmt.Println("Failed to read Silk file:", err)
// 		return
// 	}

// 	// Convert audio data to hex string
// 	hexString := hex.EncodeToString(data)
// 	// fmt.Println("Silk hexString:", hexString)

// 	url := "http://10.211.4.239:30001"
// 	err = NewWechatEncryption(url).SendVoiceMsg("wxid_3435654360314", hexString)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// }

func TestSendTextMsgNoSrc(t *testing.T) {
	url := "https://app.aworld.ltd:9112/api"
	ip := "127.0.0.1"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	err := NewWechat(ip, port, url, appSecret).SendTextMsgNoSrc("filehelper", "e1")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestStickyChat(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	err := NewWechat(ip, port, url, appSecret).StickyChat("filehelper")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUnpinChat(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	err := NewWechat(ip, port, url, appSecret).UnpinChat("filehelper")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSendPicMsg(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	err := NewWechat(ip, port, url, appSecret).SendPicMsg("filehelper", "", "1.png")
	if err != nil {
		t.Error(err)
		return
	}
}

//func TestConfigureMsgReciveFullURL(t *testing.T) {
//	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
//	err := ConfigureMsgReciveFullURL("https://mp.aworld.ltd/message/save/configure")
//	if err != nil {
//		t.Error(err)
//		return
//	}
//}

func TestConfigureMsgRecive(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	err := NewWechat(ip, port, url, appSecret).ConfigureMsgRecive(1, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestQueryBodyInfo(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	info, err := NewWechat(ip, port, url, appSecret).QueryBodyInfo("filehelper")
	if err != nil {
		t.Error(err)
		return
	}
	if info.Retstatus != "Everything is OK" {
		fmt.Println(info.Retstatus)
		return
	}
	//text, err := json.Marshal(info)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}

	fmt.Println(info.Nickname)
	fmt.Println(info.Wxid)
	fmt.Println(info.Sex)
}

func TestGetFriendOrChatroomDetailInfo(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	info, err := NewWechat(ip, port, url, appSecret).GetFriendOrChatroomDetailInfo("filehelper")
	if err != nil {
		t.Error(err)
		return
	}

	text, err := json.Marshal(info)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(text))
}

func TestGetGIFURL(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	info, err := NewWechat(ip, port, url, appSecret).GetGIFURL("<msg><emoji fromusername = \"wxid_3435654360314\" tousername = \"wxid_eow8vrunogtu41\" type=\"2\" idbuffer=\"media:0_0\" md5=\"1a73a92b7a29f751508b0b616295f275\" len = \"39592\" productid=\"com.tencent.xin.emoticon.person.stiker_155025192599bbfc37f73285f0\" androidmd5=\"1a73a92b7a29f751508b0b616295f275\" androidlen=\"39592\" s60v3md5 = \"1a73a92b7a29f751508b0b616295f275\" s60v3len=\"39592\" s60v5md5 = \"1a73a92b7a29f751508b0b616295f275\" s60v5len=\"39592\" cdnurl = \"http://wxapp.tc.qq.com/262/20304/stodownload?m=1a73a92b7a29f751508b0b616295f275&amp;filekey=30350201010421301f020201060402535a04101a73a92b7a29f751508b0b616295f2750203009aa8040d00000004627466730000000131&amp;hy=SZ&amp;storeid=32303231303632383034303932363030303433623835613737353866343466623135356630393030303030313036&amp;bizid=1023\" designerid = \"\" thumburl = \"http://mmbiz.qpic.cn/mmemoticon/ajNVdqHZLLBVWrxFJdiaibzk1KNpEoFTJoCciciaz4A3aKXYQib0Pib66bB3UYuy07Tjhs/0\" encrypturl = \"http://wxapp.tc.qq.com/262/20304/stodownload?m=c24320e18785785de22016c6facd2af1&amp;filekey=30350201010421301f020201060402535a0410c24320e18785785de22016c6facd2af10203009ab0040d00000004627466730000000131&amp;hy=SZ&amp;storeid=32303231303632383034303932363030303737393631613737353866343462386464356630393030303030313036&amp;bizid=1023\" aeskey= \"6f060f7f108a864a24257703286fa41f\" externurl = \"http://wxapp.tc.qq.com/262/20304/stodownload?m=5461c7539d727516045486953d83ebda&amp;filekey=30340201010420301e020201060402535a04105461c7539d727516045486953d83ebda02022520040d00000004627466730000000131&amp;hy=SZ&amp;storeid=32303231303632383034303932363030303934396561613737353866343430623338356630393030303030313036&amp;bizid=1023\" externmd5 = \"c2dcf907b0cd68183ce2b89115780759\" width= \"240\" height= \"240\" tpurl= \"\" tpauthkey= \"\" attachedtext= \"\" attachedtextcolor= \"\" lensid= \"\" emojiattr= \"\" linkid= \"\" desc= \"Cg8KBXpoX2NuEgbmmZrlrokKCQoFemhfdHcSAAoLCgdkZWZhdWx0EgA=\" ></emoji> <gameext type=\"0\" content=\"0\" ></gameext></msg>")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(info)
}

func TestDownPic(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"

	filePath := "D:\\test2.jpg"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	err := NewWechat(ip, port, url, appSecret).DownPic(filePath, "<msg><img hdlength=\"0\" length=\"209763\" hevc_mid_size=\"0\"  cdnbigimgurl=\"\" cdnmidimgurl=\"3057020100044b304902010002044be37a2b02032f597d0204179f3cb70204638f1b70042465613135306638322d633766322d343161302d626239352d3830363064636633613035610204011800020201000405004c53d900\" aeskey=\"f398b9444dd72d30a5a04902d232f8fb\" cdnthumburl=\"3057020100044b304902010002044be37a2b02032f597d0204179f3cb70204638f1b70042465613135306638322d633766322d343161302d626239352d3830363064636633613035610204011800020201000405004c53d900\" cdnthumblength=\"27673\" cdnthumbwidth=\"116\" cdnthumbheight=\"120\" cdnthumbaeskey=\"f398b9444dd72d30a5a04902d232f8fb\" encryver=\"1\" md5=\"12a8fd3164293a29a7592b2e4bd1a580\"/><commenturl></commenturl></msg>")
	if err != nil {
		t.Error(err)
		return
	}

	NewWechat(ip, port, url, appSecret).SendPicMsg("filehelper", filePath, "")
}

func TestGetUnReadMsgNum(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	info, err := NewWechat(ip, port, url, appSecret).GetUnReadMsgNum()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(info)
}

func TestEditFriendMark(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	err := NewWechat(ip, port, url, appSecret).EditFriendMark("", "test2")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSendDIYMsg(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	err := NewWechat(ip, port, url, appSecret).SendDIYMsg("filehelper", "3", "<?xml version=\"1.0\"?>\n<msg>\n\t<img aeskey=\"b2337cba8409dfbc309a4854c3b7a4cb\" encryver=\"0\" cdnthumbaeskey=\"b2337cba8409dfbc309a4854c3b7a4cb\" cdnthumburl=\"3057020100044b304902010002044be37a2b02032f59e1020432c7587d020463b996d0042435616561623063632d653966342d346538642d396562652d363637366263663263323935020401190a020201000405004c54a100\" cdnthumblength=\"2515\" cdnthumbheight=\"90\" cdnthumbwidth=\"120\" cdnmidheight=\"0\" cdnmidwidth=\"0\" cdnhdheight=\"0\" cdnhdwidth=\"0\" cdnmidimgurl=\"3057020100044b304902010002044be37a2b02032f59e1020432c7587d020463b996d0042435616561623063632d653966342d346538642d396562652d363637366263663263323935020401190a020201000405004c54a100\" length=\"1\" md5=\"6293dad30a9383b85ee4248308f803ab\" hevc_mid_size=\"28226\" />\n\t<platform_signature></platform_signature>\n\t<imgdatahash></imgdatahash>\n</msg>\n")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTimelineGetFristPage(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	info, err := NewWechat(ip, port, url, appSecret).TimelineGetFristPage()
	if err != nil {
		t.Error(err)
		return
	}

	text, err := json.Marshal(info)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(text))
}

func TestGetFriendTimeline(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/api", "127.0.0.1", 8100)
	ip := "10.211.4.239"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	info, err := NewWechat(ip, port, url, appSecret).GetFriendTimeline("xiaoming709959")
	if err != nil {
		t.Error(err)
		return
	}
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(info)
	// fmt.Println(bf.String())

	// text, err := json.Marshal(info)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	fmt.Println(bf.String())
}

func TestQueryDB(t *testing.T) {
	url := "https://app.aworld.ltd:9112/api"
	ip := "127.0.0.1"
	port := "30001"
	appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"
	info, err := NewWechat(ip, port, url, appSecret).QueryDB("MSG0.db", "select * from MSG limit 2")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(info)
}

// func TestProtocolBuffer(t *testing.T) {
// 	// MessageEnvelope是models.pb.go的结构体
// 	oldData := ipc.TimelineGetFristPageResult{
// 		Havenewmsg: "1",
// 	}

// 	items := make([]*ipc.TimelineGetFristPageResultDataItem, 0)

// 	item := &ipc.TimelineGetFristPageResultDataItem{
// 		MsgidFull10: "1",
// 		MsgidFull16: "2",
// 		Msgid1:      "3",
// 		Msgid2:      "4",
// 		Msgtime:     "5",
// 		Wxid:        "6",
// 		Nickname:    "7",
// 		Content:     "8",
// 	}

// 	items = append(items, item)

// 	item = &ipc.TimelineGetFristPageResultDataItem{
// 		MsgidFull10: "1a",
// 		MsgidFull16: "2b",
// 		Msgid1:      "3c",
// 		Msgid2:      "4d",
// 		Msgtime:     "5e",
// 		Wxid:        "6f",
// 		Nickname:    "7g",
// 		Content:     "8h中午",
// 	}

// 	items = append(items, item)

// 	oldData.Data = items

// 	data, err := proto.Marshal(&oldData)
// 	if err != nil {
// 		fmt.Println("marshal error: ", err.Error())
// 		return
// 	}
// 	// fmt.Println("marshal data : ", data)

// 	// text := string(data)
// 	fmt.Println("data : ", data)

// 	hexDataStr := hex.EncodeToString(data)
// 	fmt.Println("hex.EncodeToString : ", hexDataStr)

// 	// r := bytes.Runes(bArr)

// 	// // res := ""
// 	// for _, b := range r {
// 	// 	// res = fmt.Sprintf("%s%.8b", res, b)
// 	// 	fmt.Print(string(b))
// 	// }

// 	// hexDataStr := hex.EncodeToString(data)
// 	// fmt.Println("hex.EncodeToString : ", hexDataStr)

// 	// newBytes, _ := hex.DecodeString(hexDataStr)

// 	// newData2 := &ipc.TimelineGetFristPageResult{}
// 	// err = proto.Unmarshal(bytes.Runes(), newData2)
// 	// if err != nil {
// 	// 	fmt.Println("unmarshal err:", err)
// 	// }
// 	// fmt.Println("unmarshal data : ", newData2)

// 	// proto.

// }

func TestCount(t *testing.T) {

	testInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

	testInt2 := make([][]int, 0)
	testInt3 := make([]int, 0)

	for index, r := range testInt {
		// fmt.Println("(index+1)%4:", (index+1)%4)
		if (index+1)%4 != 0 {
			fmt.Println("(index+1)%4 != 0:", r)
			testInt3 = append(testInt3, r)
		} else {
			fmt.Println("(index+1)%4 == 0:", r)
			testInt3 = append(testInt3, r)
			testInt2 = append(testInt2, testInt3)
			testInt3 = make([]int, 0)
		}
		if index == len(testInt)-1 {
			testInt2 = append(testInt2, testInt3)
		}
	}

	for _, r := range testInt2 {
		fmt.Println("testInt2:", r)
	}
}

// func TestInfo(t *testing.T) {

// 	uuid := uuid.New()
// 	key := uuid.String()

// 	fmt.Println("key:", key)
// }
