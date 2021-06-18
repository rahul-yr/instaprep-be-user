package firebasedb

import (
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
)


var (
	firestoreOneInstance *firestore.Client
	firestoreOneOnce     sync.Once
)

func GetFirestoreOneInstance() *firestore.Client {
	firestoreOneOnce.Do(func() {
		conf := &firebase.Config{ProjectID: os.Getenv("APP_FIREBASE_PROJECT")}
		app, err := firebase.NewApp(global_ctx, conf)
		if err != nil {
			log.Fatalln(err)
		}
		firestoreOneInstance, err = app.Firestore(global_ctx)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("[firestore_operations] : created singleton firestoreOne instance..")
		// <-- thread safe
	})
	return firestoreOneInstance
}

func CloseFirestoreOneInstance() {
	if firestoreOneInstance != nil {
		firestoreOneInstance.Close()
	}
}