package model

type GetDBPwdResult struct {
	Key string `json:"key"`
}

type GetDBHandleAndPathResult struct {
	Data []GetDBHandleAndPathResultDataItem `json:"data"`
}
type GetDBHandleAndPathResultDataItem struct {
	Dbhandle string `json:"dbhandle"`
	Dbpath   string `json:"dbpath"`
}

type QueryDBReq struct {
	Dbname string `json:"dbname"`
	SQL    string `json:"sql"`
}

type QueryDBResult struct {
	Data []QueryDBResultDataItem `json:"data"`
}
type QueryDBResultDataItem struct {
	LocalID      string `json:"localId"`
	TalkerID     string `json:"TalkerId"`
	MsgSvrID     string `json:"MsgSvrID"`
	Type         string `json:"Type"`
	SubType      string `json:"SubType"`
	IsSender     string `json:"IsSender"`
	CreateTime   string `json:"CreateTime"`
	Sequence     string `json:"Sequence"`
	StatusEx     string `json:"StatusEx"`
	FlagEx       string `json:"FlagEx"`
	Status       string `json:"Status"`
	MsgServerSeq string `json:"MsgServerSeq"`
	MsgSequence  string `json:"MsgSequence"`
	StrTalker    string `json:"StrTalker"`
	StrContent   string `json:"StrContent"`
	Reserved0    string `json:"Reserved0"`
	BytesExtra   string `json:"BytesExtra"`
}
