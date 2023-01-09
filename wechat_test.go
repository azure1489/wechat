package wechat

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestGetFriendAndChatRoomList(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	friendAndChatRooms, err := GetFriendAndChatRoomList(url)
	if err != nil {
		t.Error(err)
		return
	}
	if len(friendAndChatRooms.Friends) > 0 {
		for _, friend := range friendAndChatRooms.Friends {
			friendStatus, err := CheckFriendStatus(url, friend.Wxid)
			if err != nil {
				t.Error(err)
				return
			}
			if friendStatus != 1 {
				log.Println("wxid:", friend.Wxid)
				log.Println("非好友 friendStatus:", friendStatus)
				log.Println("昵称:", friend.Nickname)
				log.Println("头像:", friend.Headimg)
			}
		}
	}
}

func TestSendTextMsg(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	err := SendTextMsg(url, "filehelper", "rttt")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestStickyChat(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	err := StickyChat(url, "filehelper")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUnpinChat(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	err := UnpinChat(url, "filehelper")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSendPicMsg(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	err := SendPicMsg(url, "filehelper", "", "1.png")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestConfigureMsgReciveFullURL(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30003)
	err := ConfigureMsgReciveFullURL(url, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestConfigureMsgRecive(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30003)
	err := ConfigureMsgRecive(url, 1)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestQueryBodyInfo(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	info, err := QueryBodyInfo(url, "filehelper")
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
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	info, err := GetFriendOrChatroomDetailInfo(url, "filehelper")
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
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	info, err := GetGIFURL(url, "<msg><emoji fromusername = \"wxid_3435654360314\" tousername = \"wxid_eow8vrunogtu41\" type=\"2\" idbuffer=\"media:0_0\" md5=\"1a73a92b7a29f751508b0b616295f275\" len = \"39592\" productid=\"com.tencent.xin.emoticon.person.stiker_155025192599bbfc37f73285f0\" androidmd5=\"1a73a92b7a29f751508b0b616295f275\" androidlen=\"39592\" s60v3md5 = \"1a73a92b7a29f751508b0b616295f275\" s60v3len=\"39592\" s60v5md5 = \"1a73a92b7a29f751508b0b616295f275\" s60v5len=\"39592\" cdnurl = \"http://wxapp.tc.qq.com/262/20304/stodownload?m=1a73a92b7a29f751508b0b616295f275&amp;filekey=30350201010421301f020201060402535a04101a73a92b7a29f751508b0b616295f2750203009aa8040d00000004627466730000000131&amp;hy=SZ&amp;storeid=32303231303632383034303932363030303433623835613737353866343466623135356630393030303030313036&amp;bizid=1023\" designerid = \"\" thumburl = \"http://mmbiz.qpic.cn/mmemoticon/ajNVdqHZLLBVWrxFJdiaibzk1KNpEoFTJoCciciaz4A3aKXYQib0Pib66bB3UYuy07Tjhs/0\" encrypturl = \"http://wxapp.tc.qq.com/262/20304/stodownload?m=c24320e18785785de22016c6facd2af1&amp;filekey=30350201010421301f020201060402535a0410c24320e18785785de22016c6facd2af10203009ab0040d00000004627466730000000131&amp;hy=SZ&amp;storeid=32303231303632383034303932363030303737393631613737353866343462386464356630393030303030313036&amp;bizid=1023\" aeskey= \"6f060f7f108a864a24257703286fa41f\" externurl = \"http://wxapp.tc.qq.com/262/20304/stodownload?m=5461c7539d727516045486953d83ebda&amp;filekey=30340201010420301e020201060402535a04105461c7539d727516045486953d83ebda02022520040d00000004627466730000000131&amp;hy=SZ&amp;storeid=32303231303632383034303932363030303934396561613737353866343430623338356630393030303030313036&amp;bizid=1023\" externmd5 = \"c2dcf907b0cd68183ce2b89115780759\" width= \"240\" height= \"240\" tpurl= \"\" tpauthkey= \"\" attachedtext= \"\" attachedtextcolor= \"\" lensid= \"\" emojiattr= \"\" linkid= \"\" desc= \"Cg8KBXpoX2NuEgbmmZrlrokKCQoFemhfdHcSAAoLCgdkZWZhdWx0EgA=\" ></emoji> <gameext type=\"0\" content=\"0\" ></gameext></msg>")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(info)
}

func TestDownPic(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)

	filePath := "D:\\test2.jpg"
	err := DownPic(url, filePath, "<msg><img hdlength=\"0\" length=\"209763\" hevc_mid_size=\"0\"  cdnbigimgurl=\"\" cdnmidimgurl=\"3057020100044b304902010002044be37a2b02032f597d0204179f3cb70204638f1b70042465613135306638322d633766322d343161302d626239352d3830363064636633613035610204011800020201000405004c53d900\" aeskey=\"f398b9444dd72d30a5a04902d232f8fb\" cdnthumburl=\"3057020100044b304902010002044be37a2b02032f597d0204179f3cb70204638f1b70042465613135306638322d633766322d343161302d626239352d3830363064636633613035610204011800020201000405004c53d900\" cdnthumblength=\"27673\" cdnthumbwidth=\"116\" cdnthumbheight=\"120\" cdnthumbaeskey=\"f398b9444dd72d30a5a04902d232f8fb\" encryver=\"1\" md5=\"12a8fd3164293a29a7592b2e4bd1a580\"/><commenturl></commenturl></msg>")
	if err != nil {
		t.Error(err)
		return
	}

	SendPicMsg(url, "filehelper", filePath, "")
}

func TestGetUnReadMsgNum(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	info, err := GetUnReadMsgNum(url)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(info)
}

func TestEditFriendMark(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	err := EditFriendMark(url, "", "test2")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSendDIYMsg(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "10.211.4.239", 30001)
	err := SendDIYMsg(url, "filehelper", "3", "<?xml version=\"1.0\"?>\n<msg>\n\t<img aeskey=\"b2337cba8409dfbc309a4854c3b7a4cb\" encryver=\"0\" cdnthumbaeskey=\"b2337cba8409dfbc309a4854c3b7a4cb\" cdnthumburl=\"3057020100044b304902010002044be37a2b02032f59e1020432c7587d020463b996d0042435616561623063632d653966342d346538642d396562652d363637366263663263323935020401190a020201000405004c54a100\" cdnthumblength=\"2515\" cdnthumbheight=\"90\" cdnthumbwidth=\"120\" cdnmidheight=\"0\" cdnmidwidth=\"0\" cdnhdheight=\"0\" cdnhdwidth=\"0\" cdnmidimgurl=\"3057020100044b304902010002044be37a2b02032f59e1020432c7587d020463b996d0042435616561623063632d653966342d346538642d396562652d363637366263663263323935020401190a020201000405004c54a100\" length=\"1\" md5=\"6293dad30a9383b85ee4248308f803ab\" hevc_mid_size=\"28226\" />\n\t<platform_signature></platform_signature>\n\t<imgdatahash></imgdatahash>\n</msg>\n")
	if err != nil {
		t.Error(err)
		return
	}
}
