package login

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rahul-yr/instaprep-be-user/auth"
	pubsubgen "github.com/rahul-yr/instaprep-be-user/auth/pubsub_gen"
	redisdb "github.com/rahul-yr/instaprep-be-user/auth/redis_auth_gen_db"
)

const (
	// example : OTP::admin@email.com
	otp_store_key = "OTP::"
)

func Getkey(email string) string {
	return otp_store_key + email
}

func SendOTP(c *fiber.Ctx) error {
	// extract request params
	requestParams := new(RequestParams)
	if err := c.BodyParser(requestParams); err != nil {
		return c.Status(404).JSON(&Response{Message: "Please verify inputs", Status: false})
	}
	// validate input params
	err := requestParams.Validate()
	if err != nil {
		return c.Status(404).JSON(&Response{Message: err.Error(), Status: false})
	}
	// send otp
	status := sendingOTP(requestParams.Email)
	if !status {
		return c.Status(404).JSON(&Response{Message: "Failed to sent OTP", Status: false})
	}
	return c.Status(200).JSON(&Response{Message: "Email has been sent successfully", Status: true})
}

func sendingOTP(email string) bool {
	// declare variables
	red := &redisdb.RedisOneOps{}
	tokenGenService := &auth.TokenGeneratorService{}
	// create otp
	otp_generated, err := tokenGenService.GenerateOTPToken()
	if err != nil {
		// failed to create otp so return false
		return false
	}
	// store otp in redis
	//  15 mins OTP Valid
	expiry_time := 15 * time.Minute
	err = red.StoreKV(Getkey(email), otp_generated, expiry_time)
	if err != nil {
		// failed to store generated otp in redis, so returning false
		log.Printf("redis : Error %+v \n", err)
		return false
	}
	// ingest to pub sub topic with (email, otp_generated)
	payload_obj := &PubsubEmailMessage{
		Type: "otp",
		Payload: map[string]interface{}{
			"email": email,
			"otp":   otp_generated,
		},
	}
	email_service := &pubsubgen.EmailEvent{}
	err = email_service.Publish(payload_obj)
	if err != nil {
		log.Printf("pubsub : Error %+v \n", err)
	}
	// return true if no error else false
	return err == nil
}
