package pubsubgen

import (
	"log"
	"sync"
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/pubsub"
)


var (
	emailTopicInstance *pubsub.Topic
	emailSyncOnce sync.Once
	emailTopicName   = "email-notifications-test-topic"
)

type EmailEvent struct{
	// TODO
}

// Creates and returns singleton email topic instance
func GetEmailTopicInstance() *pubsub.Topic {
	emailSyncOnce.Do(func() {
		pubsub_client := GetPubsubOneInstance()
		emailTopicInstance = pubsub_client.Topic(emailTopicName)
		log.Println("[email_topic_instance] : created singleton pubsub email topic instance..")
	})
	return emailTopicInstance
}

func(helpers *EmailEvent) Publish(data interface{}) error {
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	email_pubsub_client := GetEmailTopicInstance()
	byte_map, err := json.Marshal(data)
	if err != nil {
		return err
	}
	result_obj := email_pubsub_client.Publish(timeout, &pubsub.Message{
		Data: byte_map,
	})
	_, err = result_obj.Get(timeout)
	if err != nil {
		return err
	}
	return nil
}