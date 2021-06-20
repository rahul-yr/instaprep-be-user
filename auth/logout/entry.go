package logout

import (
	"github.com/gofiber/fiber/v2"
)

const (
	jwt_token_key = "TOKEN::"
)

func Getkey(email string, tk string) string {
	return jwt_token_key + email + "::" + tk
}

func Logout(c *fiber.Ctx) error {
	// extract request params
	// requestParams := new(RequestParams)
	// if err := c.BodyParser(requestParams); err != nil {
	// 	return c.Status(404).JSON(&Response{Status: false})
	// }

	// validate token
	// tokenGenService := &auth.TokenGeneratorService{}
	// tokenMapper, err := tokenGenService.VerifyAuthToken(requestParams.Token)
	// if err != nil {
	// 	return c.Status(404).JSON(&Response{Status: false})
	// }
	// delete token from redis
	// red := &redisdb.RedisOneOps{}
	// err = red.DeleteKey(Getkey(tokenMapper.Email, requestParams.Token))
	// if err != nil {
	// 	return c.Status(404).JSON(&Response{Status: false})
	// }
	return c.Status(200).JSON(&Response{Status: true})
}
