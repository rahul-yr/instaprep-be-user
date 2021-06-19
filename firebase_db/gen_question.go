package firebasedb

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
)

// type Question struct{
	// TODO
// }


// Create, Read, Update, Delete
const (
	ERROR_QUESTION_CREATION = "question : failed to create question details"
	ERROR_QUESTION_UPDATE   = "question : failed to update question details"
	ERROR_QUESTION_DELETE   = "question : failed to delete question"
)

var questionCollectionRef *firestore.CollectionRef
var questionSyncOnce sync.Once

// Creates and returns singleton question collection reference
func GetQuestionCollectionReference() *firestore.CollectionRef {
	questionSyncOnce.Do(func() {
		questionCollectionRef = GetFirestoreOneInstance().Collection(QUESTIONS_COLLECTION)
	})
	return questionCollectionRef
}

// Creates a question with given object details
func (v *Question) Create() (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, ws, err := GetQuestionCollectionReference().Add(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_QUESTION_CREATION)
	}
	return docRef.ID, nil
}

// Creates a question along with ID with given object details
func (v *Question) CreateWithIDField() error {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetQuestionCollectionReference().NewDoc()
	v.ID = docRef.ID
	ws, err := docRef.Set(timeout, v)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New(ERROR_QUESTION_CREATION)
	}
	return nil
}

func (v *Question) CreateWithId(hash_id string) (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetQuestionCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_QUESTION_CREATION)
	}
	return hash_id, nil
}

// returns a given question details if exists
func (v *Question) Read(hash_id string) (*Question, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, err := GetQuestionCollectionReference().Doc(hash_id).Get(timeout)
	if err != nil {
		return nil, err
	}
	r := &Question{}
	docRef.DataTo(r)
	return r, nil
}

// returns list of question details if exists
func (v *Question) ReadMultipleIds(hash_id []string) ([]*Question, error) {
	itemsList := make([]*Question,0)
	for _, i := range hash_id {
		item, err := v.Read(i)
		if err != nil{
			return nil, err
		}
		itemsList = append(itemsList, item)
	}
	return itemsList, nil
}

// returns a list of question details if exists
func (v *Question) ReadAll() ([]*Question, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docSnap, err := GetQuestionCollectionReference().Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*Question, 0)
	for _, item := range docSnap {
		r := &Question{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// returns a list of question details if exists
func (v *Question) ReadAllByCondition(fields []string, operators []string, val []interface{}) ([]*Question, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetQuestionCollectionReference().Query
	for i := 0; i < len(fields); i++ {
		docRef = docRef.Where(fields[i], operators[i], val[i])
	}

	docSnap, err := docRef.Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*Question, 0)
	for _, item := range docSnap {
		r := &Question{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// Update entire question document with given new question object
//
// retuns the update status
func (v *Question) Update(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetQuestionCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_QUESTION_UPDATE)
	}
	return true, nil
}

// This deletes the question object itself
func (v *Question) Delete(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetQuestionCollectionReference().Doc(hash_id).Delete(timeout)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_QUESTION_DELETE)
	}
	return true, nil
}
