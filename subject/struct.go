package subject

import firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"

type Response struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TopicIds    []string `json:"topic_ids"`
}

type RequestParams struct {
	LearningPathId string `json:"learning_path_id"  xml:"learning_path_id" form:"learning_path_id"`
}

func (v *Response) GetResponseObject(item *firebasedb.Subject) *Response {
	temp := &Response{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		TopicIds:    item.TopicIds,
	}
	return temp
}

func (v *Response) GetResponseObjectList(item []*firebasedb.Subject) []*Response {
	result := make([]*Response, 0)
	for _, item := range item {
		temp := &Response{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			TopicIds:    item.TopicIds,
		}
		result = append(result, temp)
	}
	return result
}
