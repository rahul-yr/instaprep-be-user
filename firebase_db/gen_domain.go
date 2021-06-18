package firebasedb

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
)

// type Domain struct{
	// TODO
// }


// Create, Read, Update, Delete
const (
	ERROR_DOMAIN_CREATION = "domain : failed to create domain details"
	ERROR_DOMAIN_UPDATE   = "domain : failed to update domain details"
	ERROR_DOMAIN_DELETE   = "domain : failed to delete domain"
)

var domainCollectionRef *firestore.CollectionRef
var domainSyncOnce sync.Once

// Creates and returns singleton domain collection reference
func GetDomainCollectionReference() *firestore.CollectionRef {
	domainSyncOnce.Do(func() {
		domainCollectionRef = GetFirestoreOneInstance().Collection(DOMAINS_COLLECTION)
	})
	return domainCollectionRef
}

// Creates a domain with given object details
func (v *Domain) Create() (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, ws, err := GetDomainCollectionReference().Add(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_DOMAIN_CREATION)
	}
	return docRef.ID, nil
}

// Creates a domain along with ID with given object details
func (v *Domain) CreateWithIDField() error {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetDomainCollectionReference().NewDoc()
	v.ID = docRef.ID
	ws, err := docRef.Set(timeout, v)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New(ERROR_DOMAIN_CREATION)
	}
	return nil
}

func (v *Domain) CreateWithId(hash_id string) (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetDomainCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_DOMAIN_CREATION)
	}
	return hash_id, nil
}

// returns a given domain details if exists
func (v *Domain) Read(hash_id string) (*Domain, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, err := GetDomainCollectionReference().Doc(hash_id).Get(timeout)
	if err != nil {
		return nil, err
	}
	r := &Domain{}
	docRef.DataTo(r)
	return r, nil
}

// returns list of domain details if exists
func (v *Domain) ReadMultipleIds(hash_id ...string) ([]*Domain, error) {
	itemsList := make([]*Domain,0)
	for _, i := range hash_id {
		item, err := v.Read(i)
		if err != nil{
			return nil, err
		}
		itemsList = append(itemsList, item)
	}
	return itemsList, nil
}

// returns a list of domain details if exists
func (v *Domain) ReadAll() ([]*Domain, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docSnap, err := GetDomainCollectionReference().Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*Domain, 0)
	for _, item := range docSnap {
		r := &Domain{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// returns a list of domain details if exists
func (v *Domain) ReadAllByCondition(fields []string, operators []string, val []interface{}) ([]*Domain, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetDomainCollectionReference().Query
	for i := 0; i < len(fields); i++ {
		docRef = docRef.Where(fields[i], operators[i], val[i])
	}

	docSnap, err := docRef.Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*Domain, 0)
	for _, item := range docSnap {
		r := &Domain{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// Update entire domain document with given new domain object
//
// retuns the update status
func (v *Domain) Update(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetDomainCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_DOMAIN_UPDATE)
	}
	return true, nil
}

// This deletes the domain object itself
func (v *Domain) Delete(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetDomainCollectionReference().Doc(hash_id).Delete(timeout)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_DOMAIN_DELETE)
	}
	return true, nil
}
