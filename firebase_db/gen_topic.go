package firebasedb

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
)

// type Topic struct{
	// TODO
// }


// Create, Read, Update, Delete
const (
	ERROR_TOPIC_CREATION = "topic : failed to create topic details"
	ERROR_TOPIC_UPDATE   = "topic : failed to update topic details"
	ERROR_TOPIC_DELETE   = "topic : failed to delete topic"
)

var topicCollectionRef *firestore.CollectionRef
var topicSyncOnce sync.Once

// Creates and returns singleton topic collection reference
func GetTopicCollectionReference() *firestore.CollectionRef {
	topicSyncOnce.Do(func() {
		topicCollectionRef = GetFirestoreOneInstance().Collection(TOPICS_COLLECTION)
	})
	return topicCollectionRef
}

// Creates a topic with given object details
func (v *Topic) Create() (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, ws, err := GetTopicCollectionReference().Add(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_TOPIC_CREATION)
	}
	return docRef.ID, nil
}

// Creates a topic along with ID with given object details
func (v *Topic) CreateWithIDField() error {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetTopicCollectionReference().NewDoc()
	v.ID = docRef.ID
	ws, err := docRef.Set(timeout, v)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New(ERROR_TOPIC_CREATION)
	}
	return nil
}

func (v *Topic) CreateWithId(hash_id string) (string, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetTopicCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return "", err
	}
	if ws == nil {
		return "", errors.New(ERROR_TOPIC_CREATION)
	}
	return hash_id, nil
}

// returns a given topic details if exists
func (v *Topic) Read(hash_id string) (*Topic, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef, err := GetTopicCollectionReference().Doc(hash_id).Get(timeout)
	if err != nil {
		return nil, err
	}
	r := &Topic{}
	docRef.DataTo(r)
	return r, nil
}

// returns list of topic details if exists
func (v *Topic) ReadMultipleIds(hash_id []string) ([]*Topic, error) {
	itemsList := make([]*Topic,0)
	for _, i := range hash_id {
		item, err := v.Read(i)
		if err != nil{
			return nil, err
		}
		itemsList = append(itemsList, item)
	}
	return itemsList, nil
}

// returns a list of topic details if exists
func (v *Topic) ReadAll() ([]*Topic, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docSnap, err := GetTopicCollectionReference().Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*Topic, 0)
	for _, item := range docSnap {
		r := &Topic{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// returns a list of topic details if exists
func (v *Topic) ReadAllByCondition(fields []string, operators []string, val []interface{}) ([]*Topic, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	docRef := GetTopicCollectionReference().Query
	for i := 0; i < len(fields); i++ {
		docRef = docRef.Where(fields[i], operators[i], val[i])
	}

	docSnap, err := docRef.Documents(timeout).GetAll()
	if err != nil {
		return nil, err
	}
	output := make([]*Topic, 0)
	for _, item := range docSnap {
		r := &Topic{}
		item.DataTo(r)
		output = append(output, r)
	}
	return output, nil
}


// Update entire topic document with given new topic object
//
// retuns the update status
func (v *Topic) Update(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetTopicCollectionReference().Doc(hash_id).Set(timeout, v)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_TOPIC_UPDATE)
	}
	return true, nil
}

// This deletes the topic object itself
func (v *Topic) Delete(hash_id string) (bool, error) {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	ws, err := GetTopicCollectionReference().Doc(hash_id).Delete(timeout)
	if err != nil {
		return false, err
	}
	if ws == nil {
		return false, errors.New(ERROR_TOPIC_DELETE)
	}
	return true, nil
}
