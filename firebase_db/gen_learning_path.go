package firebasedb

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
)

// type LearningPath struct{
	// TODO
// }


// Create, Read, Update, Delete
const (
	ERROR_LEARNING_PATH_CREATION = "learningPath : failed to create learningPath details"
	ERROR_LEARNING_PATH_UPDATE   = "learningPath : failed to update learningPath details"
	ERROR_LEARNING_PATH_DELETE   = "learningPath : failed to delete learningPath"
)

var learningPathCollectionRef *firestore.CollectionRef
var learningPathSyncOnce sync.Once

// Creates and returns singleton learningPath collection reference
func GetLearningPathCollectionReference() *firestore.CollectionRef {
	learningPathSyncOnce.Do(func() {
		learningPathCollectionRef = GetFirestoreOneInstance().Collection(LEARNING_PATHS_COLLECTION)
	})
	return learningPathCollectionRef
}

// Creates a learningPath with given object details
func (v *LearningPath) Create() (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, ws, err := GetLearningPathCollectionReference().Add(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_LEARNING_PATH_CREATION)
	}
	return docRef.ID, nil
}

// Creates a learningPath along with ID with given object details
func (v *LearningPath) CreateWithIDField() error {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetLearningPathCollectionReference().NewDoc()
	v.ID = docRef.ID
	ws, err := docRef.Set(timeout, v)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New(ERROR_LEARNING_PATH_CREATION)
	}
	return nil
}

func (v *LearningPath) CreateWithId(hash_id string) (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetLearningPathCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_DOMAIN_CREATION)
	}
	return hash_id, nil
}

// returns a given learningPath details if exists
func (v *LearningPath) Read(hash_id string) (*LearningPath, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, err := GetLearningPathCollectionReference().Doc(hash_id).Get(timeout)
	if err != nil {
		return nil, err
	}
	r := &LearningPath{}
	docRef.DataTo(r)
	return r, nil
}

// returns list of learningPath details if exists
func (v *LearningPath) ReadMultipleIds(hash_id ...string) ([]*LearningPath, error) {
	itemsList := make([]*LearningPath,0)
	for _, i := range hash_id {
		item, err := v.Read(i)
		if err != nil{
			return nil, err
		}
		itemsList = append(itemsList, item)
	}
	return itemsList, nil
}

// returns a list of learningPath details if exists
func (v *LearningPath) ReadAll() ([]*LearningPath, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docSnap, err := GetLearningPathCollectionReference().Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*LearningPath, 0)
	for _, item := range docSnap {
		r := &LearningPath{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// returns a list of learningPath details if exists
func (v *LearningPath) ReadAllByCondition(fields []string, operators []string, val []interface{}) ([]*LearningPath, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetLearningPathCollectionReference().Query
	for i := 0; i < len(fields); i++ {
		docRef = docRef.Where(fields[i], operators[i], val[i])
	}

	docSnap, err := docRef.Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*LearningPath, 0)
	for _, item := range docSnap {
		r := &LearningPath{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// Update entire learningPath document with given new learningPath object
//
// retuns the update status
func (v *LearningPath) Update(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetLearningPathCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_LEARNING_PATH_UPDATE)
	}
	return true, nil
}

// This deletes the learningPath object itself
func (v *LearningPath) Delete(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetLearningPathCollectionReference().Doc(hash_id).Delete(timeout)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_LEARNING_PATH_DELETE)
	}
	return true, nil
}
