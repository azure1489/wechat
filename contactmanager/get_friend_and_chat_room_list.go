package contactmanager

import (
	"encoding/json"
)

type GetFriendAndChatRoomListResult struct {
	Friend   []Friend   `json:"friend"`
	Chatroom []Chatroom `json:"chatroom"`
	Gh       []Gh       `json:"gh"`
	Openim   []Openim   `json:"openim"`
}

// {"index":"514","wxid":"l304255890","account":"z11280325","nickname":"钟","v3":"v3_020b3826fd03010000000000949c081914e053000000501","markname":"","starrole":"3","dontseeit":"1","dontseeme":"1","lag":""}
type Friend struct {
	Index     string `json:"index"`
	Wxid      string `json:"wxid"`
	Account   string `json:"account"`
	Markname  string `json:"markname"`
	Nickname  string `json:"nickname"`
	V3        string `json:"v3"`
	Starrole  string `json:"starrole"`
	Dontseeit string `json:"dontseeit"`
	Dontseeme string `json:"dontseeme"`
	Lag       string `json:"lag"`
}

// {"index":"34","gid":"18936197505@chatroom","gname":"财务系统开发","markname":"","v3":""},
type Chatroom struct {
	Index    string `json:"index"`
	Gid      string `json:"gid"`
	Markname string `json:"markname"`
	Gname    string `json:"gname"`
	V3       string `json:"v3"`
}

// {"index":"439","wxid":"gh_33c67e91b4c7","account":"salud-yogurt","nickname":"Salud撒露冻酸奶","markname":"","v3":""}
type Gh struct {
	Index    string `json:"index"`
	Wxid     string `json:"wxid"`
	Account  string `json:"account"`
	Markname string `json:"markname"`
	Nickname string `json:"nickname"`
	V3       string `json:"v3"`
}

// {"index":"1","wxid":"25984983121544334@openim","account":"","nickname":"","v3":"","markname":"","starrole":"4","dontseeit":"0","dontseeme":"0"}
type Openim struct {
	Index     string `json:"index"`
	Wxid      string `json:"wxid"`
	Account   string `json:"account"`
	Nickname  string `json:"nickname"`
	V3        string `json:"v3"`
	Markname  string `json:"markname"`
	Starrole  string `json:"starrole"`
	Dontseeit string `json:"dontseeit"`
	Dontseeme string `json:"dontseeme"`
}

type GetFriendAndChatRoomListReq struct {
	FriendType string `json:"type"`
}

func (c *ContactManagerServiceImpl) GetFriendAndChatRoomList(friendType string) (*GetFriendAndChatRoomListResult, error) {

	req := GetFriendAndChatRoomListReq{
		FriendType: friendType,
	}

	resultBody, err := c.http.DoPost("/GetFriendAndChatRoomList", req)
	if err != nil {
		return nil, err
	}

	// fmt.Println("GetFriendAndChatRoomList resultBody:", string(resultBody))

	commonResult := GetFriendAndChatRoomListResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
