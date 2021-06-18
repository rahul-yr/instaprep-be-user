package domain

import firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"

type Response struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (v *Response) GetResponseObject(item *firebasedb.Domain) *Response {
	temp := &Response{
		Name: item.Name,
		ID:   item.ID,
	}
	return temp
}

func (v *Response) GetResponseObjectList(item []*firebasedb.Domain) []*Response {
	result := make([]*Response, 0)
	for _, one := range item {
		temp := &Response{Name: one.Name, ID: one.ID}
		result = append(result, temp)
	}
	return result
}
