package topic

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"
	"github.com/rahul-yr/instaprep-be-user/helpers"
	redisdb "github.com/rahul-yr/instaprep-be-user/redis_db"
)

const (
	all_topics_key = "user:all_topic:"
	// in hours
	redis_expiry_time = 48
)

func Getkey(subject_id string) string {
	return all_topics_key + subject_id
}

func GetTopicsBySubject(c *fiber.Ctx) error {
	// extract request params
	requestParams := new(RequestParams)
	if err := c.BodyParser(requestParams); err != nil {
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Please verify inputs", Status: false})
	}

	// variables
	red := &redisdb.RedisOneOps{}
	fire := &firebasedb.Topic{}
	fire_sub := &firebasedb.Subject{}

	// get cache
	var cachedList []*Response
	if err := red.GetJSON(Getkey(requestParams.SubjectId), &cachedList); err == nil {
		return c.Status(200).JSON(cachedList)
	}
	// fire
	// get the learning_path details from firestore
	// extract the details of subject_ids
	lp_obj, err := fire_sub.Read(requestParams.SubjectId)
	if err != nil {
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}
	// here obj holds the list of subject ids
	topic_ids := lp_obj.TopicIds

	all_docs, err := fire.ReadMultipleIds(topic_ids)
	if err != nil {
		log.Printf("error : %+v \n", err)
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}

	// store cache
	if err = red.StoreJSON(Getkey(requestParams.SubjectId), all_docs, redis_expiry_time*time.Hour); err != nil {
		log.Printf("error : %+v \n", err)
	}

	var temp *Response
	res := temp.GetResponseObjectList(all_docs)
	return c.Status(200).JSON(res)
}
