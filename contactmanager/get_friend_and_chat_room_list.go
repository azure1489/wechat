package contactmanager

import "encoding/json"

type GetFriendAndChatRoomListResult struct {
	Friend   []Friend   `json:"friend"`
	Chatroom []Chatroom `json:"chatroom"`
	Gh       []Gh       `json:"gh"`
}

type Friend struct {
	Index     string `json:"index"`
	Wxid      string `json:"wxid"`
	Account   string `json:"account"`
	Markname  string `json:"markname"`
	Nickname  string `json:"nickname"`
	Headimg   string `json:"headimg"`
	V3        string `json:"v3"`
	Sex       string `json:"sex"`
	Starrole  string `json:"starrole"`
	Dontseeit string `json:"dontseeit"`
	Dontseeme string `json:"dontseeme"`
	Lag       string `json:"lag"`
}

type Chatroom struct {
	Index    string `json:"index"`
	Gid      string `json:"gid"`
	Markname string `json:"markname"`
	Gname    string `json:"gname"`
	V3       string `json:"v3"`
}

type Gh struct {
	Index    string `json:"index"`
	Wxid     string `json:"wxid"`
	Account  string `json:"account"`
	Markname string `json:"markname"`
	Nickname string `json:"nickname"`
	Headimg  string `json:"headimg"`
	V3       string `json:"v3"`
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

	commonResult := GetFriendAndChatRoomListResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
