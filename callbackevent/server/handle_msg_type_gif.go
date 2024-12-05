package server

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"

	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeGif 处理自定义表情消息
func (srv *Server) handleMsgTypeGif(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	// 自定义表情消息 "msgtype":"47",
	gifMsg := message.Gif{
		GifPath: msgInfo.Get("gif_path").String(),
	}

	msgContent := wcMsgItem.Msg
	fromType := wcMsgItem.FromType

	if emojiMsg, err := getEmojiMsg(msgContent); err == nil {

		gifMsg.CdnURL = emojiMsg.Emoji.CdnURL
		base64Desc := emojiMsg.Emoji.Desc
		if base64Desc != "" {
			// fmt.Println("base64Desc:", base64Desc)
			if desc, err := decodeString(base64Desc); err == nil {
				// fmt.Println("decodeString()  desc:", desc)
				desc = strings.TrimSpace(desc)
				desc = strings.ReplaceAll(desc, "\n", "")
				// fmt.Println("decodeString() - TrimSpace  ReplaceAll desc:", desc)
				if desc2, err := getDesc(desc); err == nil {
					desc2 = strings.Replace(desc2, "default", "", 1)
					gifMsg.Desc = desc2
					// fmt.Println("getDesc() desc2:", desc)
				}

			}
		}
	}

	wcMsgItem.MsgItem = gifMsg

	// fromtype: 1=个人消息,2=群消息
	if fromType == "1" {
		wcMsgItem.EventType = message.PCRecvGifImgMsgEvent
	} else if fromType == "2" {
		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
			FromGname: msgInfo.Get("fromgname").String(), // 群名称
			FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
		}
		wcMsgItem.EventType = message.PCRecvGroupGifImgMsgEvent
	}
	return nil
}

// 过去Emoji消息
func getEmojiMsg(emojiXmlStr string) (*message.EmojiMsg, error) {

	var msg message.EmojiMsg
	err := xml.Unmarshal([]byte(emojiXmlStr), &msg)
	if err != nil {
		fmt.Println("Error unmarshalling XML: ", err)
		return nil, err
	}

	// fmt.Printf("Parsed Struct: %+v\n", msg)

	return &msg, nil
}

// 解码 Base64 字符串
func decodeString(base64String string) (string, error) {

	// 解码 Base64 字符串
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Error decoding string: ", err.Error())
		return "", err
	}

	return string(data), nil
}

func getDesc(desc string) (string, error) {

	// 正则表达式匹配 'zh_cn' 后的任意字符，直到 'zh_tw' 之前
	re := regexp.MustCompile(`zh_cn\s*(.*?)\s*zh_tw`)

	match := re.FindStringSubmatch(desc)
	// && len(match) > 1
	if match != nil {
		// 去除换行符和额外的空格
		cleanedText := strings.ReplaceAll(match[1], "\n", "")
		cleanedText = strings.TrimSpace(cleanedText)
		cleanedText = compressStr(cleanedText)

		// fmt.Println("Matched Text:", cleanedText)

		return cleanedText, nil

	} else {
		// fmt.Println("No match found")

		return desc, nil
	}
}

// 利用正则表达式压缩字符串，去除空格或制表符
func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, "")
}
