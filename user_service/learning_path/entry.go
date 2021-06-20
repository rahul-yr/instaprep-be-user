package learningpath

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"
	"github.com/rahul-yr/instaprep-be-user/helpers"
	redisdb "github.com/rahul-yr/instaprep-be-user/redis_db"
)

const (
	all_learning_paths_key = "user:all_learning_path"
	// in hours
	redis_expiry_time = 48
)

func Getkey() string {
	return all_learning_paths_key
}

func GetAllLearningPathByDomain(c *fiber.Ctx) error {
	// extract request params
	// variables
	red := &redisdb.RedisOneOps{}
	fire := &firebasedb.LearningPath{}

	// get cache
	var cachedList []*Response
	if err := red.GetJSON(Getkey(), &cachedList); err == nil {
		return c.Status(200).JSON(&fiber.Map{"results": cachedList, "status": true})
	}

	// fire
	all_docs, err := fire.ReadAllByCondition([]string{"published"}, []string{"=="}, []interface{}{true})
	if err != nil {
		log.Printf("error : %+v \n", err)
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}
	if len(all_docs) == 0 {
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}
	// store cache
	if err = red.StoreJSON(Getkey(), all_docs, redis_expiry_time*time.Hour); err != nil {
		log.Printf("error : %+v \n", err)
	}

	var temp *Response
	res := temp.GetResponseObjectList(all_docs)
	return c.Status(200).JSON(&fiber.Map{"results": res, "status": true})
}
