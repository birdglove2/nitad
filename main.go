package main

import (
	"os"

	"github.com/birdglove2/nitad-backend/api"
	"github.com/birdglove2/nitad-backend/api/project"
	"github.com/birdglove2/nitad-backend/config"
	"github.com/birdglove2/nitad-backend/cronjob"
	"github.com/birdglove2/nitad-backend/database"
	"github.com/birdglove2/nitad-backend/errors"
	"github.com/birdglove2/nitad-backend/gcp"
	"github.com/birdglove2/nitad-backend/redis"
	"github.com/birdglove2/nitad-backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"go.uber.org/zap"
)

var PORT = os.Getenv("PORT")

func main() {
	utils.InitZap()

	config.Loadenv()
	envErr := config.Checkenv()
	if envErr != nil {
		zap.S().Warn(envErr.Error())
		os.Exit(1)
	}

	database.ConnectDb(os.Getenv("MONGO_URI"))
	defer database.DisconnectDb()

	uploader := gcp.Init()
	redisService := redis.Init()

	app := config.InitApp()
	app.Use(cache.New(cache.Config{
		Expiration: redis.DefaultCacheExpireTime,
		Storage:    redisService.GetStorage(),
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Path() + "?" + string(c.Request().URI().QueryString())
		},
		Next: func(c *fiber.Ctx) bool {
			isTrue := project.IsGetProjectPath(c) // handle incrementing view in cache
			return isTrue
		},
	}))

	api.CreateAPI(app, uploader, redisService)

	cronjob.Init(redisService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "result": "Hello, this is NITAD Backend Server v2.2  !"})
	})

	app.All("*", func(c *fiber.Ctx) error {
		return errors.Throw(c, errors.NewNotFoundError("Page"))
	})

	if PORT == "" {
		PORT = "3000"
	}

	zap.S().Info("===== Running on ", os.Getenv("APP_ENV"), " stage =====")
	zap.S().Info("===== Listening to port ", PORT, " ======")

	err := app.Listen(":" + PORT)
	if err != nil {
		zap.S().Warn("Listen to " + PORT + " Failed!")
		zap.S().Warn("Error: ", err.Error())
	}
}
