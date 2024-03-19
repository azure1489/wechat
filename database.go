package wechat

import (
	"encoding/json"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// GetDBPwd 查询微信本地数据库密钥 https://www.showdoc.com.cn/WeChatProject/8962025006749060
func (w *Wechat) GetDBPwd() (*model.GetDBPwdResult, error) {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return nil, err
	}

	resultBody, err := client.DoPost("/GetDBPwd", nil)
	if err != nil {
		return nil, err
	}

	commonResult := model.GetDBPwdResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}

// GetDBHandleAndPath 取数据库操作句柄和文件名 https://www.showdoc.com.cn/WeChatProject/8984969039725273
func (w *Wechat) GetDBHandleAndPath() (*[]model.GetDBHandleAndPathResultDataItem, error) {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return nil, err
	}

	resultBody, err := client.DoPost("/GetDBHandleAndPath", nil)
	if err != nil {
		return nil, err
	}

	commonResult := model.GetDBHandleAndPathResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult.Data, nil
}

// QueryDB 取数据库操作句柄和文件名 https://www.showdoc.com.cn/WeChatProject/8986208243971483
func (w *Wechat) QueryDB(dbname, sql string) (string, error) {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return "", err
	}

	req := model.QueryDBReq{
		Dbname: dbname,
		SQL:    sql,
	}

	resultBody, err := client.DoPost("/QueryDB", req)
	if err != nil {
		return "", err
	}

	// commonResult := model.QueryDBResult{}
	// err = json.Unmarshal(resultBody, &commonResult)
	// if err != nil {
	// 	return "", err
	// }

	return string(resultBody), nil
}

// QueryDB 取数据库操作句柄和文件名 https://www.showdoc.com.cn/WeChatProject/9159818632076714
func (w *Wechat) QueryDBAllTable(dbname, sql string) (string, error) {
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.PublicKeyPath, 0)
	if err != nil {
		return "", err
	}

	req := model.QueryDBReq{
		Dbname: dbname,
		SQL:    sql,
	}

	resultBody, err := client.DoPost("/QueryDB", req)
	if err != nil {
		return "", err
	}

	// commonResult := model.QueryDBResult{}
	// err = json.Unmarshal(resultBody, &commonResult)
	// if err != nil {
	// 	return "", err
	// }

	return string(resultBody), nil
}
