package learningpath

import firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"

type Response struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	DomainId         string   `json:"domain_id"`
	QuestionLevelIds []string `json:"question_level_ids"`
}

type RequestParams struct {
	DomainId string `json:"domain_id" xml:"domain_id" form:"domain_id"`
}

func (v *Response) GetResponseObject(item *firebasedb.LearningPath) *Response {
	temp := &Response{
		ID:               item.ID,
		Name:             item.Name,
		Description:      item.Description,
		DomainId:         item.DomainId,
		QuestionLevelIds: item.QuestionLevelIds,
	}
	return temp
}

func (v *Response) GetResponseObjectList(item []*firebasedb.LearningPath) []*Response {
	result := make([]*Response, 0)
	for _, item := range item {
		temp := &Response{
			ID:               item.ID,
			Name:             item.Name,
			Description:      item.Description,
			DomainId:         item.DomainId,
			QuestionLevelIds: item.QuestionLevelIds,
		}
		result = append(result, temp)
	}
	return result
}
