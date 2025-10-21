package bootstrap

import (
	"fmt"
	"time"

	"github.com/bccfilkom-be/postify/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func StartApp() error {
	conf := config.Load()
	app := fiber.New(fiber.Config{
		// ErrorHandler: FiberErrorHandler,
	})

	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:               50,
		Expiration:        60 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests. Please try again later.",
			})
		},
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Your API is Running!",
		})
	})

	addr := fmt.Sprintf("localhost:%s", conf.AppPort)
	if conf.AppEnv == "production" {
		addr = fmt.Sprintf("0.0.0.0:%s", conf.AppPort)
	}

	return app.Listen(addr)
}
