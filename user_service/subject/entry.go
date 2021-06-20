package subject

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"
	"github.com/rahul-yr/instaprep-be-user/helpers"
	redisdb "github.com/rahul-yr/instaprep-be-user/redis_db"
)

const (
	all_subjects_key = "user:all_subject:"
	// in hours
	redis_expiry_time = 48
)

func Getkey(learning_path_id string) string {
	return all_subjects_key + learning_path_id
}

func GetSubjectsByLearningPath(c *fiber.Ctx) error {
	// extract request params
	requestParams := new(RequestParams)
	if err := c.BodyParser(requestParams); err != nil {
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Please verify inputs", Status: false})
	}

	// variables
	red := &redisdb.RedisOneOps{}
	fire := &firebasedb.Subject{}
	fire_lp := &firebasedb.LearningPath{}

	// get cache
	var cachedList []*Response
	if err := red.GetJSON(Getkey(requestParams.LearningPathId), &cachedList); err == nil {
		return c.Status(200).JSON(&fiber.Map{"results": cachedList, "status": true})
	}
	// fire
	// get the learning_path details from firestore
	// extract the details of subject_ids
	lp_obj, err := fire_lp.Read(requestParams.LearningPathId)
	if err != nil {
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}
	// if learning_path is inactive
	if !lp_obj.Published {
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}
	// here obj holds the list of subject ids
	subject_ids := lp_obj.SubjectIds

	all_docs, err := fire.ReadMultipleIds(subject_ids)
	if err != nil {
		log.Printf("error : %+v \n", err)
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}

	// store cache
	if err = red.StoreJSON(Getkey(requestParams.LearningPathId), all_docs, redis_expiry_time*time.Hour); err != nil {
		log.Printf("error : %+v \n", err)
	}

	var temp *Response
	res := temp.GetResponseObjectList(all_docs)
	return c.Status(200).JSON(&fiber.Map{"results": res, "status": true})
}
