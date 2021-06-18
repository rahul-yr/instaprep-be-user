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
	all_learnings_key = "user:all_learning_path:"
	// in hours
	redis_expiry_time = 48
)

func Getkey(domain string) string {
	return all_learnings_key + domain
}

func GetAllLearningPathByDomain(c *fiber.Ctx) error {
	// extract request params
	requestParams := new(RequestParams)
	if err := c.BodyParser(requestParams); err != nil {
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Please verify inputs", Status: false})
	}
	// variables
	red := &redisdb.RedisOneOps{}
	fire := &firebasedb.LearningPath{}

	// get cache
	var cachedList []*Response
	if err := red.GetJSON(Getkey(requestParams.DomainId), &cachedList); err == nil {
		return c.Status(200).JSON(cachedList)
	}

	// fire
	all_docs, err := fire.ReadAllByCondition([]string{"published", "id"}, []string{"==", "=="}, []interface{}{true, requestParams.DomainId})
	if err != nil {
		log.Printf("error : %+v \n", err)
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}
	if len(all_docs) == 0 {
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}
	// store cache
	if err = red.StoreJSON(Getkey(requestParams.DomainId), all_docs, redis_expiry_time*time.Hour); err != nil {
		log.Printf("error : %+v \n", err)
	}

	var temp *Response
	res := temp.GetResponseObjectList(all_docs)
	return c.Status(200).JSON(res)
}
