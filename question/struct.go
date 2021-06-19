package question

import firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"

type PageCountResponse struct {
	Results    []*Response `json:"results"`
	TotalPages int         `json:"total_pages"`
}

type Response struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	OptionA  string `json:"a"`
	OptionB  string `json:"b"`
	OptionC  string `json:"c"`
	OptionD  string `json:"d"`
	Answer   string `json:"answer"`
	Level    string `json:"level_id"`
}

type RequestParams struct {
	TopicId string `json:"topic_id"  xml:"topic_id" form:"topic_id"`
	PageNum int    `json:"page_num"  xml:"page_num" form:"page_num"`
}

func (r *RequestParams) VerifyPageNum() {
	if r.PageNum == 0 {
		r.PageNum = 1
	}
}

func (v *Response) GetResponseObject(item *firebasedb.Question) *Response {
	temp := &Response{
		ID:       item.ID,
		Question: item.Question,
		OptionA:  item.OptionA,
		OptionB:  item.OptionB,
		OptionC:  item.OptionC,
		OptionD:  item.OptionD,
		Answer:   item.Answer,
		Level:    item.Level,
	}
	return temp
}

func (v *Response) GetResponseObjectList(item []*firebasedb.Question) []*Response {
	result := make([]*Response, 0)
	for _, item := range item {
		temp := &Response{
			ID:       item.ID,
			Question: item.Question,
			OptionA:  item.OptionA,
			OptionB:  item.OptionB,
			OptionC:  item.OptionC,
			OptionD:  item.OptionD,
			Answer:   item.Answer,
			Level:    item.Level,
		}
		result = append(result, temp)
	}
	return result
}

func (v *Response) GetPageCountResonseObject(item []*firebasedb.Question, total_pages int) *PageCountResponse {
	res := v.GetResponseObjectList(item)
	return &PageCountResponse{Results: res, TotalPages: total_pages}
}
