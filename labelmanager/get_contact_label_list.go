package labelmanager

import "encoding/json"

//	{
//	    "inc": "3",
//	    "data": [
//	        {
//	            "name": "万达",
//	            "id": "1"
//	        },
//	        {
//	            "name": "金科",
//	            "id": "2"
//	        },
//	        {
//	            "name": "碧桂园",
//	            "id": "3"
//	        }
//	    ]
//	}
type GetContactLabelListResult struct {
	Inc  string `json:"inc"`
	Data []struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"data"`
}

func (l *LabelManagerServiceImpl) GetContactLabelList() (*GetContactLabelListResult, error) {
	resultBody, err := l.http.DoPost("/GetContactLabelList", nil)
	if err != nil {
		return nil, err
	}

	commonResult := GetContactLabelListResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
