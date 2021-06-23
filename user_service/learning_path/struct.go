package learningpath

import firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"

type Response struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

func (v *Response) GetResponseObject(item *firebasedb.LearningPath) *Response {
	temp := &Response{
		ID:          item.ID,
		Name:        item.Name,
		Icon:        item.Icon,
		Description: item.Description,
	}
	return temp
}

func (v *Response) GetResponseObjectList(item []*firebasedb.LearningPath) []*Response {
	result := make([]*Response, 0)
	for _, item := range item {
		temp := &Response{
			ID:          item.ID,
			Name:        item.Name,
			Icon:        item.Icon,
			Description: item.Description,
		}
		result = append(result, temp)
	}
	return result
}
