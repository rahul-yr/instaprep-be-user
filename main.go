package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	learningpath "github.com/rahul-yr/instaprep-be-user/user_service/learning_path"
	question "github.com/rahul-yr/instaprep-be-user/user_service/question"
	questionlevel "github.com/rahul-yr/instaprep-be-user/user_service/question_level"
	subject "github.com/rahul-yr/instaprep-be-user/user_service/subject"
	topic "github.com/rahul-yr/instaprep-be-user/user_service/topic"

	login "github.com/rahul-yr/instaprep-be-user/auth/login"
	logout "github.com/rahul-yr/instaprep-be-user/auth/logout"
	refreshtoken "github.com/rahul-yr/instaprep-be-user/auth/refresh_token"
	verifyotp "github.com/rahul-yr/instaprep-be-user/auth/verify_otp"
	verifytoken "github.com/rahul-yr/instaprep-be-user/auth/verify_token"
)

func init() {
	godotenv.Load(".env")

}

func main() {
	// fiber api for multi cores
	app := fiber.New(fiber.Config{Prefork: true})
	setupSecurityConfigs(app)
	// Init routes

	api_auth := app.Group("/auth")
	api_user := app.Group("/user")

	// authentication middleware
	api_user.Use(func(c *fiber.Ctx) error {
		authorization_header := string(c.Request().Header.Peek("Authorization"))
		actual_token := strings.Split(authorization_header, " ")
		if len(actual_token) == 2 {
			// verify authentication
			status := verifytoken.VerifyTokenMiddleware(actual_token[1])
			if status {
				return c.Next()
			}
		}
		return c.Status(403).JSON(&fiber.Map{"status": "false"})
	})

	// routes
	setupAuthRoutes(api_auth)
	setupRoutes(api_user)
	// listen
	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))

}

func setupAuthRoutes(app fiber.Router) {
	// This method send otp based on given email id
	//
	// @inputs	>>	email
	app.Post("/login", login.SendOTP)

	// This method returns jwt token based on successfull otp validation
	//
	// @inputs	>>	email, otp
	app.Post("/verify-otp", verifyotp.VerifyOTP)

	// This method returns true on successfull validation of token
	//
	// @inputs	>>	token
	app.Post("/verify-token", verifytoken.VerifyToken)

	// This method returns new token if the provided token is valid and about to expiry
	//
	// @inputs	>>	token
	app.Post("/refresh-token", refreshtoken.RefreshToken)

	// This method logouts the existing user from redis db
	// for now we won't call this method
	//
	// @inputs	>>	token
	app.Post("/logout", logout.Logout)

}

func setupRoutes(app fiber.Router) {
	// Get all Available Question levels
	//
	// no input params required
	app.Post("/question-level/all", questionlevel.GetAllPracticeTestTypes)

	// Get all LearningPath
	//
	// no input params required
	app.Post("/learning-path/all", learningpath.GetAllLearningPathByDomain)

	// Get all Subjects based on learning_path_id
	//
	// @inputs	>>	learning_path_id
	app.Post("/subject/all", subject.GetSubjectsByLearningPath)

	// Get all Topics based on subject_id
	//
	// @inputs	>>	subject_id
	app.Post("/topic/all", topic.GetTopicsBySubject)

	// Get All Questions based on topic_id,  page_num
	//
	// @inputs	>>	topic_id,page_num
	//
	app.Post("/question/all", question.GetQuestionsByTopic)
}

func setupSecurityConfigs(app *fiber.App) {
	// Prod Configs
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	app.Use(compress.New())
	app.Use(cors.New())
	// end of prod configs
}
