package question

import (
	"log"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"
	"github.com/rahul-yr/instaprep-be-user/helpers"
	redisdb "github.com/rahul-yr/instaprep-be-user/redis_db"
)

const (
	all_questions_key = "user:all_question:"
	// in hours
	redis_expiry_time    = 48
	questions_page_count = 30
)

func GetkeyWithPageId(topic_id string, page_num int) string {
	// sample questions format >> user:all_question:topic_id:page:page_id
	//
	// user:all_question:${topic_id}:page:${page_id}
	return all_questions_key + topic_id + ":page:" + strconv.Itoa(page_num)
}
func GetTotalPageCountKey(topic_id string) string {
	// user:all_question:${topic_id}:total_page_count
	return all_questions_key + topic_id + ":total_page_count"
}

func GetQuestionsByTopic(c *fiber.Ctx) error {
	// extract request params
	requestParams := new(RequestParams)
	if err := c.BodyParser(requestParams); err != nil {
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Please verify inputs", Status: false})
	}
	// required validation methods
	requestParams.VerifyPageNum()

	// decalre variables
	red := &redisdb.RedisOneOps{}
	fire := &firebasedb.Question{}

	// cache response if available
	var pageCountRes *PageCountResponse
	if err := red.GetJSON(GetkeyWithPageId(requestParams.TopicId, requestParams.PageNum), pageCountRes); err == nil {
		return c.Status(200).JSON(&fiber.Map{"results": pageCountRes, "status": true})
	}

	//	if error get total page count from redis
	// if total page count is available then we have cache
	// one of 2 scenarios i.e 1. invalid page number(means out of bounds) or
	// some internal redis client request context failure
	total_page_count, err := red.GetKV(GetTotalPageCountKey(requestParams.TopicId))
	if err == nil {
		tpc, err := strconv.Atoi(total_page_count)
		if err != nil {
			return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
		}
		// if requested page count is out of bounds return error
		if tpc < requestParams.PageNum {
			// means requested page count is beyond the available pages count
			return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
		}
	}
	// if total_page count key is not available in redis or requested page count is within limits but still error
	// recalculate the data again
	// get all questions from firestore related to that topic_id
	all_docs, err := fire.ReadAllByCondition([]string{"topic_id"}, []string{"=="}, []interface{}{requestParams.TopicId})
	if err != nil {
		log.Printf("error : %+v \n", err)
		return c.Status(404).JSON(&helpers.ErrorResponse{Error: "Something went wrong", Status: false})
	}
	// find the count and create all the redis cache items
	numberOfQuestions := len(all_docs)
	numberOfPages := int(math.Ceil((float64(numberOfQuestions) / questions_page_count)))
	// if calculated page count is 0, then make it atleast 1 in reality
	if numberOfPages == 0 {
		numberOfPages = 1
	}
	// declare variables
	var pcr PageCountResponse
	var temp *Response
	var arraySlice []*firebasedb.Question

	for i := 0; i < numberOfPages; i++ {
		if ((i + 1) * 30) > numberOfQuestions {
			arraySlice = all_docs[i*30:]
		} else {
			arraySlice = all_docs[i*30 : (i+1)*30]
		}
		pageResObj := temp.GetPageCountResonseObject(arraySlice, numberOfPages)
		if requestParams.PageNum == i+1 {
			pcr = *pageResObj
		}
		// store slice of questions to redis
		err = red.StoreJSON(GetkeyWithPageId(requestParams.TopicId, i+1), pageResObj, redis_expiry_time*time.Hour)
		if err != nil {
			log.Printf("error : %+v \n", err)
		}
	}

	return c.Status(200).JSON(&fiber.Map{"results": pcr, "status": true})

}
