package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	domain "github.com/rahul-yr/instaprep-be-user/domain"
	learningpath "github.com/rahul-yr/instaprep-be-user/learning_path"
	practicetesttype "github.com/rahul-yr/instaprep-be-user/practice_test_type"
	questionlevel "github.com/rahul-yr/instaprep-be-user/question_level"
	subject "github.com/rahul-yr/instaprep-be-user/subject"
)

func init() {
	godotenv.Load(".env")

}

func main() {
	// fiber api for multi cores
	app := fiber.New(fiber.Config{Prefork: true})
	setupSecurityConfigs(app)
	// Init routes
	setupRoutes(app)
	// listen
	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))

}

func setupRoutes(app *fiber.App) {
	// Get all Available domains
	// 
	// no input params required
	app.Post("/all-domains", domain.GetAllDomains)
	
	//	Get all Available Test types allowed 
	// 
	// no input params required
	app.Post("/all-practice-test-type", practicetesttype.GetAllPracticeTestTypes)

	// Get all Available Question levels
	// 
	// no input params required
	app.Post("/all-question-level", questionlevel.GetAllPracticeTestTypes)

	// Get all LearningPath based on domain_id
	// 
	// @inputs	>>  domain_id 
	app.Post("/all-learning-path", learningpath.GetAllLearningPathByDomain)

	// Get all Subjects based on learning_path_id
	// 
	// @inputs	>>	learning_path_id	
	app.Post("/all-subject", subject.GetSubjectsByLearningPath)
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
