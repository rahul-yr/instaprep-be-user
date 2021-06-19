package practicetesttype

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"
	"github.com/rahul-yr/instaprep-be-user/helpers"
	redisdb "github.com/rahul-yr/instaprep-be-user/redis_db"
)

const (
	all_practice_test_types_key = "user:all_practice_test_type"
	// in hours
	redis_expiry_time = 48
)

func Getkey() string {
	return all_practice_test_types_key
}

// returns all domains
//
// either from cache or remote
func GetAllPracticeTestTypes(c *fiber.Ctx) error {
	// variables
	red := &redisdb.RedisOneOps{}
	fire := &firebasedb.PracticeTestType{}

	// redis check
	var cachedList []*Response
	if err := red.GetJSON(Getkey(), &cachedList); err == nil {
		return c.Status(200).JSON(cachedList)
	}

	// firestore check
	all_docs, err := fire.ReadAll()
	if err != nil {
		log.Printf("error : %+v \n", err)
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}

	// cache
	if err = red.StoreJSON(Getkey(), all_docs, redis_expiry_time*time.Hour); err != nil {
		log.Printf("error : %+v \n", err)
	}

	var temp *Response
	res := temp.GetResponseObjectList(all_docs)
	return c.Status(200).JSON(res)

}
