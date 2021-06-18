package firebasedb

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
)

// type PracticeTestType struct{
	// TODO
// }


// Create, Read, Update, Delete
const (
	ERROR_PRACTICE_TEST_TYPE_CREATION = "practiceTestType : failed to create practiceTestType details"
	ERROR_PRACTICE_TEST_TYPE_UPDATE   = "practiceTestType : failed to update practiceTestType details"
	ERROR_PRACTICE_TEST_TYPE_DELETE   = "practiceTestType : failed to delete practiceTestType"
)

var practiceTestTypeCollectionRef *firestore.CollectionRef
var practiceTestTypeSyncOnce sync.Once

// Creates and returns singleton practiceTestType collection reference
func GetPracticeTestTypeCollectionReference() *firestore.CollectionRef {
	practiceTestTypeSyncOnce.Do(func() {
		practiceTestTypeCollectionRef = GetFirestoreOneInstance().Collection(PRACTICE_TEST_TYPES_COLLECTION)
	})
	return practiceTestTypeCollectionRef
}

// Creates a practiceTestType with given object details
func (v *PracticeTestType) Create() (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, ws, err := GetPracticeTestTypeCollectionReference().Add(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_PRACTICE_TEST_TYPE_CREATION)
	}
	return docRef.ID, nil
}

// Creates a practiceTestType along with ID with given object details
func (v *PracticeTestType) CreateWithIDField() error {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetPracticeTestTypeCollectionReference().NewDoc()
	v.ID = docRef.ID
	ws, err := docRef.Set(timeout, v)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New(ERROR_PRACTICE_TEST_TYPE_CREATION)
	}
	return nil
}

func (v *PracticeTestType) CreateWithId(hash_id string) (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetPracticeTestTypeCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_DOMAIN_CREATION)
	}
	return hash_id, nil
}

// returns a given practiceTestType details if exists
func (v *PracticeTestType) Read(hash_id string) (*PracticeTestType, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, err := GetPracticeTestTypeCollectionReference().Doc(hash_id).Get(timeout)
	if err != nil {
		return nil, err
	}
	r := &PracticeTestType{}
	docRef.DataTo(r)
	return r, nil
}

// returns list of practiceTestType details if exists
func (v *PracticeTestType) ReadMultipleIds(hash_id ...string) ([]*PracticeTestType, error) {
	itemsList := make([]*PracticeTestType,0)
	for _, i := range hash_id {
		item, err := v.Read(i)
		if err != nil{
			return nil, err
		}
		itemsList = append(itemsList, item)
	}
	return itemsList, nil
}

// returns a list of practiceTestType details if exists
func (v *PracticeTestType) ReadAll() ([]*PracticeTestType, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docSnap, err := GetPracticeTestTypeCollectionReference().Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*PracticeTestType, 0)
	for _, item := range docSnap {
		r := &PracticeTestType{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// returns a list of practiceTestType details if exists
func (v *PracticeTestType) ReadAllByCondition(fields []string, operators []string, val []interface{}) ([]*PracticeTestType, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetPracticeTestTypeCollectionReference().Query
	for i := 0; i < len(fields); i++ {
		docRef = docRef.Where(fields[i], operators[i], val[i])
	}

	docSnap, err := docRef.Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*PracticeTestType, 0)
	for _, item := range docSnap {
		r := &PracticeTestType{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// Update entire practiceTestType document with given new practiceTestType object
//
// retuns the update status
func (v *PracticeTestType) Update(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetPracticeTestTypeCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_PRACTICE_TEST_TYPE_UPDATE)
	}
	return true, nil
}

// This deletes the practiceTestType object itself
func (v *PracticeTestType) Delete(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetPracticeTestTypeCollectionReference().Doc(hash_id).Delete(timeout)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_PRACTICE_TEST_TYPE_DELETE)
	}
	return true, nil
}
