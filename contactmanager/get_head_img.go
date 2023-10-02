package contactmanager

import "encoding/json"

type GetHeadIMGReq struct {
	Wxid string `json:"wxid"`
}

type GetHeadIMGResult struct {
	ImgLen string `json:"img_len"`
	ImgHex string `json:"img_hex"`
}

// GetHeadIMG 获取头像
func (c *ContactManagerServiceImpl) GetHeadIMG(wxid string) (result *GetHeadIMGResult, err error) {

	req := GetHeadIMGReq{
		Wxid: wxid,
	}

	resultBody, err := c.http.DoPost("/GetHeadImg", req)
	if err != nil {
		return nil, err
	}

	commonResult := GetHeadIMGResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
