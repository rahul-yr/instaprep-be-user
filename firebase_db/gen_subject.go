package firebasedb

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
)

// type Subject struct{
	// TODO
// }


// Create, Read, Update, Delete
const (
	ERROR_SUBJECT_CREATION = "subject : failed to create subject details"
	ERROR_SUBJECT_UPDATE   = "subject : failed to update subject details"
	ERROR_SUBJECT_DELETE   = "subject : failed to delete subject"
)

var subjectCollectionRef *firestore.CollectionRef
var subjectSyncOnce sync.Once

// Creates and returns singleton subject collection reference
func GetSubjectCollectionReference() *firestore.CollectionRef {
	subjectSyncOnce.Do(func() {
		subjectCollectionRef = GetFirestoreOneInstance().Collection(SUBJECTS_COLLECTION)
	})
	return subjectCollectionRef
}

// Creates a subject with given object details
func (v *Subject) Create() (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, ws, err := GetSubjectCollectionReference().Add(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_SUBJECT_CREATION)
	}
	return docRef.ID, nil
}

// Creates a subject along with ID with given object details
func (v *Subject) CreateWithIDField() error {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetSubjectCollectionReference().NewDoc()
	v.ID = docRef.ID
	ws, err := docRef.Set(timeout, v)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New(ERROR_SUBJECT_CREATION)
	}
	return nil
}

func (v *Subject) CreateWithId(hash_id string) (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetSubjectCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_SUBJECT_CREATION)
	}
	return hash_id, nil
}

// returns a given subject details if exists
func (v *Subject) Read(hash_id string) (*Subject, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, err := GetSubjectCollectionReference().Doc(hash_id).Get(timeout)
	if err != nil {
		return nil, err
	}
	r := &Subject{}
	docRef.DataTo(r)
	return r, nil
}

// returns list of subject details if exists
func (v *Subject) ReadMultipleIds(hash_id []string) ([]*Subject, error) {
	itemsList := make([]*Subject,0)
	for _, i := range hash_id {
		item, err := v.Read(i)
		if err != nil{
			return nil, err
		}
		itemsList = append(itemsList, item)
	}
	return itemsList, nil
}

// returns a list of subject details if exists
func (v *Subject) ReadAll() ([]*Subject, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docSnap, err := GetSubjectCollectionReference().Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*Subject, 0)
	for _, item := range docSnap {
		r := &Subject{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// returns a list of subject details if exists
func (v *Subject) ReadAllByCondition(fields []string, operators []string, val []interface{}) ([]*Subject, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetSubjectCollectionReference().Query
	for i := 0; i < len(fields); i++ {
		docRef = docRef.Where(fields[i], operators[i], val[i])
	}

	docSnap, err := docRef.Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*Subject, 0)
	for _, item := range docSnap {
		r := &Subject{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// Update entire subject document with given new subject object
//
// retuns the update status
func (v *Subject) Update(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetSubjectCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_SUBJECT_UPDATE)
	}
	return true, nil
}

// This deletes the subject object itself
func (v *Subject) Delete(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetSubjectCollectionReference().Doc(hash_id).Delete(timeout)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_SUBJECT_DELETE)
	}
	return true, nil
}
