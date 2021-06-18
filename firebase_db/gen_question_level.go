package firebasedb

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
)

// type QuestionLevel struct{
	// TODO
// }


// Create, Read, Update, Delete
const (
	ERROR_QUESTION_LEVEL_CREATION = "questionLevel : failed to create questionLevel details"
	ERROR_QUESTION_LEVEL_UPDATE   = "questionLevel : failed to update questionLevel details"
	ERROR_QUESTION_LEVEL_DELETE   = "questionLevel : failed to delete questionLevel"
)

var questionLevelCollectionRef *firestore.CollectionRef
var questionLevelSyncOnce sync.Once

// Creates and returns singleton questionLevel collection reference
func GetQuestionLevelCollectionReference() *firestore.CollectionRef {
	questionLevelSyncOnce.Do(func() {
		questionLevelCollectionRef = GetFirestoreOneInstance().Collection(QUESTION_LEVELS_COLLECTION)
	})
	return questionLevelCollectionRef
}

// Creates a questionLevel with given object details
func (v *QuestionLevel) Create() (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, ws, err := GetQuestionLevelCollectionReference().Add(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_QUESTION_LEVEL_CREATION)
	}
	return docRef.ID, nil
}

// Creates a questionLevel along with ID with given object details
func (v *QuestionLevel) CreateWithIDField() error {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetQuestionLevelCollectionReference().NewDoc()
	v.ID = docRef.ID
	ws, err := docRef.Set(timeout, v)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New(ERROR_QUESTION_LEVEL_CREATION)
	}
	return nil
}

func (v *QuestionLevel) CreateWithId(hash_id string) (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetQuestionLevelCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_DOMAIN_CREATION)
	}
	return hash_id, nil
}

// returns a given questionLevel details if exists
func (v *QuestionLevel) Read(hash_id string) (*QuestionLevel, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, err := GetQuestionLevelCollectionReference().Doc(hash_id).Get(timeout)
	if err != nil {
		return nil, err
	}
	r := &QuestionLevel{}
	docRef.DataTo(r)
	return r, nil
}

// returns list of questionLevel details if exists
func (v *QuestionLevel) ReadMultipleIds(hash_id ...string) ([]*QuestionLevel, error) {
	itemsList := make([]*QuestionLevel,0)
	for _, i := range hash_id {
		item, err := v.Read(i)
		if err != nil{
			return nil, err
		}
		itemsList = append(itemsList, item)
	}
	return itemsList, nil
}

// returns a list of questionLevel details if exists
func (v *QuestionLevel) ReadAll() ([]*QuestionLevel, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docSnap, err := GetQuestionLevelCollectionReference().Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*QuestionLevel, 0)
	for _, item := range docSnap {
		r := &QuestionLevel{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// returns a list of questionLevel details if exists
func (v *QuestionLevel) ReadAllByCondition(fields []string, operators []string, val []interface{}) ([]*QuestionLevel, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetQuestionLevelCollectionReference().Query
	for i := 0; i < len(fields); i++ {
		docRef = docRef.Where(fields[i], operators[i], val[i])
	}

	docSnap, err := docRef.Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*QuestionLevel, 0)
	for _, item := range docSnap {
		r := &QuestionLevel{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// Update entire questionLevel document with given new questionLevel object
//
// retuns the update status
func (v *QuestionLevel) Update(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetQuestionLevelCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_QUESTION_LEVEL_UPDATE)
	}
	return true, nil
}

// This deletes the questionLevel object itself
func (v *QuestionLevel) Delete(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetQuestionLevelCollectionReference().Doc(hash_id).Delete(timeout)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_QUESTION_LEVEL_DELETE)
	}
	return true, nil
}
