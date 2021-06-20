package refreshtoken

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahul-yr/instaprep-be-user/auth"
)

const (
	jwt_token_key = "TOKEN::"
)

func Getkey(email string, tk string) string {
	return jwt_token_key + email + "::" + tk
}

func RefreshToken(c *fiber.Ctx) error {
	// extract request params
	requestParams := new(RequestParams)
	if err := c.BodyParser(requestParams); err != nil {
		return c.Status(404).JSON(&Response{Status: false})
	}

	// validate token
	tokenGenService := &auth.TokenGeneratorService{}
	new_token, err := tokenGenService.RefreshAuthToken(requestParams.Token)
	if err != nil {
		return c.Status(404).JSON(&Response{Status: false})
	}
	// store token to redis
	// currently commented
	err = storejwtTokenToRedis(new_token.Email, new_token.Token)
	if err != nil {
		return c.Status(404).JSON(&Response{Status: false})
	}
	return c.Status(200).JSON(&Response{Token: new_token.Token, Status: true})
}

func storejwtTokenToRedis(email string, token string) error {
	// Uncomment below code for token consistency using redis
	// declare variables
	// red := &redisdb.RedisOneOps{}
	// expiry_time := 72 * time.Hour

	// err := red.StoreKV(Getkey(email, token), "on", expiry_time)
	// if err != nil {
	// 	return err
	// }
	return nil
}
