package verifytoken

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

func VerifyToken(c *fiber.Ctx) error {
	// extract request params
	requestParams := new(RequestParams)
	if err := c.BodyParser(requestParams); err != nil {
		return c.Status(404).JSON(&Response{Status: false})
	}

	// validate token
	tokenGenService := &auth.TokenGeneratorService{}
	_, err := tokenGenService.VerifyAuthToken(requestParams.Token)
	if err != nil {
		return c.Status(404).JSON(&Response{Status: false})
	}
	// Uncomment below code for token consistency using redis
	// cache validate token
	// red := &redisdb.RedisOneOps{}
	// _, err = red.GetKV(Getkey(tokenMapper.Email, requestParams.Token))
	// if err != nil {
	// 	return c.Status(404).JSON(&Response{Status: false})
	// }

	return c.Status(200).JSON(&Response{Status: true})
}

func VerifyTokenMiddleware(token string) bool {
	// validate token
	tokenGenService := &auth.TokenGeneratorService{}
	_, err := tokenGenService.VerifyAuthToken(token)
	return err == nil
}
