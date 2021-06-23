package verifyotp

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rahul-yr/instaprep-be-user/auth"
	redisdb "github.com/rahul-yr/instaprep-be-user/auth/redis_auth_gen_db"
	firebasedb "github.com/rahul-yr/instaprep-be-user/firebase_db"
)

const (
	// example : OTP::admin@email.com
	redis_key     = "OTP::"
	jwt_token_key = "TOKEN::"
)

func Getkey(email string) string {
	return redis_key + email
}

func VerifyOTP(c *fiber.Ctx) error {
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
	// declare variables
	red := &redisdb.RedisOneOps{}

	// cache
	cache_otp, err := red.GetKV(Getkey(requestParams.Email))
	if err != nil {
		return c.Status(404).JSON(&Response{Message: "Please verify input details", Status: false})
	}

	// if otp is matching
	if requestParams.OTP == cache_otp {
		// get user id from firebase
		uid, err := firebaseAccount(requestParams.Email)
		if err != nil {
			return c.Status(404).JSON(&Response{Message: "Something went wrong", Status: false})
		}
		// generate jwt token
		signed_token, err := generateToken(requestParams.Email, uid)
		if err != nil {
			return c.Status(404).JSON(&Response{Message: "Something went wrong", Status: false})
		}
		// delete otp from redis
		err = red.DeleteKey(Getkey(requestParams.Email))
		if err != nil {
			return c.Status(404).JSON(&Response{Message: "Something went wrong", Status: false})
		}
		// TODO storejwtTokenToRedis
		err = storejwtTokenToRedis(requestParams.Email, signed_token)
		if err != nil {
			return c.Status(404).JSON(&Response{Message: "Something went wrong", Status: false})
		}
		// finally return the valid jwt token
		return c.Status(200).JSON(&Response{Token: signed_token, Status: true})
	}
	// if otp is not matching
	return c.Status(404).JSON(&Response{Message: "Please verify input details", Status: false})

}

func storejwtTokenToRedis(email string, token string) error {
	// Uncomment below code for token consistency using redis
	// declare variables
	// red := &redisdb.RedisOneOps{}
	// expiry_time := 72 * time.Hour

	// err := red.StoreKV(format_jwt(email, token), "on", expiry_time)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func format_jwt(email string, tk string) string {
	return jwt_token_key + email + "::" + tk
}

func firebaseAccount(email string) (string, error) {
	fire := &firebasedb.User{}
	users_list, err := fire.ReadAllByCondition([]string{"email"}, []string{"=="}, []interface{}{email})
	if err != nil {
		return "", err
	}
	// exceptional case which wouldn't occur
	// just added as a check
	if len(users_list) > 1 {
		log.Printf("firebase_account_user_error : It seems multiple same user account exists in database %+v \n", users_list)
		return "", errors.New("user : Mutliple account exists for same user")
	}
	if len(users_list) == 0 {
		// means account not exits
		// a new user
		// create an account
		new_user := &firebasedb.User{
			Name:      "",
			Email:     email,
			Active:    true,
			FCMTokens: []string{},
			UpdatedAt: time.Now(),
		}
		err := new_user.CreateWithIDField()
		if err != nil {
			return "", err
		}
		return new_user.ID, nil

	} else {
		// existing user
		u := users_list[0]
		return u.ID, nil
	}

}

func generateToken(email string, user_id string) (string, error) {
	tokenGenService := &auth.TokenGeneratorService{}
	token, err := tokenGenService.GenerateAuthToken(email, user_id)
	if err != nil {
		return "", err
	}
	return token, nil
}
