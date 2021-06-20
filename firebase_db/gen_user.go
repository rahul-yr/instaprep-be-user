package firebasedb

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
)

// type User struct{
	// TODO
// }


// Create, Read, Update, Delete
const (
	ERROR_USER_CREATION = "user : failed to create user details"
	ERROR_USER_UPDATE   = "user : failed to update user details"
	ERROR_USER_DELETE   = "user : failed to delete user"
)

var userCollectionRef *firestore.CollectionRef
var userSyncOnce sync.Once

// Creates and returns singleton user collection reference
func GetUserCollectionReference() *firestore.CollectionRef {
	userSyncOnce.Do(func() {
		userCollectionRef = GetFirestoreOneInstance().Collection(USERS_COLLECTION)
	})
	return userCollectionRef
}

// Creates a user with given object details
func (v *User) Create() (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, ws, err := GetUserCollectionReference().Add(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_USER_CREATION)
	}
	return docRef.ID, nil
}

// Creates a user along with ID with given object details
func (v *User) CreateWithIDField() error {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetUserCollectionReference().NewDoc()
	v.ID = docRef.ID
	ws, err := docRef.Set(timeout, v)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New(ERROR_USER_CREATION)
	}
	return nil
}

func (v *User) CreateWithId(hash_id string) (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetUserCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_USER_CREATION)
	}
	return hash_id, nil
}

// returns a given user details if exists
func (v *User) Read(hash_id string) (*User, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, err := GetUserCollectionReference().Doc(hash_id).Get(timeout)
	if err != nil {
		return nil, err
	}
	r := &User{}
	docRef.DataTo(r)
	return r, nil
}

// returns list of user details if exists
func (v *User) ReadMultipleIds(hash_id []string) ([]*User, error) {
	itemsList := make([]*User,0)
	for _, i := range hash_id {
		item, err := v.Read(i)
		if err != nil{
			return nil, err
		}
		itemsList = append(itemsList, item)
	}
	return itemsList, nil
}

// returns a list of user details if exists
func (v *User) ReadAll() ([]*User, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docSnap, err := GetUserCollectionReference().Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*User, 0)
	for _, item := range docSnap {
		r := &User{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// returns a list of user details if exists
func (v *User) ReadAllByCondition(fields []string, operators []string, val []interface{}) ([]*User, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetUserCollectionReference().Query
	for i := 0; i < len(fields); i++ {
		docRef = docRef.Where(fields[i], operators[i], val[i])
	}

	docSnap, err := docRef.Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*User, 0)
	for _, item := range docSnap {
		r := &User{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// Update entire user document with given new user object
//
// retuns the update status
func (v *User) Update(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetUserCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_USER_UPDATE)
	}
	return true, nil
}

// This deletes the user object itself
func (v *User) Delete(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetUserCollectionReference().Doc(hash_id).Delete(timeout)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_USER_DELETE)
	}
	return true, nil
}
