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
	err := SendTextMsg(url, "filehelper", "中午还吃吗")
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
	err := DownPic(url, "D:\\test.jpg", "<msg><img hdlength=\"0\" length=\"315516\" hevc_mid_size=\"0\"  cdnbigimgurl=\"\" cdnmidimgurl=\"3057020100044b304902010002044be37a2b02032f59e1020415c7587d020463b70dd4042461663562393166302d313236332d346666632d393633302d3162313061313766623030640204011800020201000405004c550500\" aeskey=\"0de40e3aaef6cd046e57415849621ed8\" cdnthumburl=\"3057020100044b304902010002044be37a2b02032f59e1020415c7587d020463b70dd4042461663562393166302d313236332d346666632d393633302d3162313061313766623030640204011800020201000405004c550500\" cdnthumblength=\"16497\" cdnthumbwidth=\"120\" cdnthumbheight=\"120\" cdnthumbaeskey=\"0de40e3aaef6cd046e57415849621ed8\" encryver=\"1\" md5=\"286b5947152e8698e1caa807c1239698\"/><commenturl></commenturl></msg>")
	if err != nil {
		t.Error(err)
		return
	}
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
