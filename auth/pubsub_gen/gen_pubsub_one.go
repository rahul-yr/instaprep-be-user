package pubsubgen

import (
	"log"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
)

var (
	pubsubOneInstance *pubsub.Client
	pubsubOneOnce     sync.Once
)

func GetPubsubOneInstance() *pubsub.Client {
	pubsubOneOnce.Do(func() {
		var err error
		pubsubOneInstance, err = pubsub.NewClient(global_ctx, os.Getenv("APP_FIREBASE_PROJECT"))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("[pubsub_operations] : created singleton pubsubOne instance..")
		// <-- thread safe
	})
	return pubsubOneInstance
}

func ClosePubsubOneInstance() {
	if pubsubOneInstance != nil {
		pubsubOneInstance.Close()
	}
}