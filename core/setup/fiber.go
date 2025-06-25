package setup

import (
	"embed"
	"him/fiber-api/core/handlers"
	"him/fiber-api/core/services"
	"net/http"
	"os"
	"time"

	core_structs "him/fiber-api/core/structs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/phuslu/log"
	fiberlog "github.com/phuslu/log/fiber"
)

func SetupFiber(embedFile embed.FS) *fiber.App {
	fiberEngine := fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHandler,
	})

	fiberEngine.Use(func(c *fiber.Ctx) error {
		appState := &core_structs.AppState{
			Validator: validator.New(),
		}
		c.Locals("state", appState)
		return c.Next()
	})

	SetupLogger(fiberEngine)
	SetupCompress(fiberEngine)
	SetupLimiter(fiberEngine)
	SetupCORS(fiberEngine)

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add EncryptCookie Middleware")
	}
	fiberEngine.Use(encryptcookie.New(encryptcookie.Config{
		Key: os.Getenv("SECRET_KEY"),
	}))

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add Favicon Middleware")
	}
	fiberEngine.Use(favicon.New(favicon.Config{
		// File: "./favicon.ico",
		// URL:  "/favicon.ico",
		FileSystem: http.FS(embedFile),
	}))

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add ETag Middleware")
	}
	fiberEngine.Use(etag.New())

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add Helmet Middleware")
	}
	fiberEngine.Use(helmet.New())

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add RequestID Middleware")
	}
	fiberEngine.Use(requestid.New())

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add Recover Middleware")
	}
	fiberEngine.Use(recover.New())

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add Skip Middleware")
	}
	fiberEngine.Use(skip.New(func(ctx *fiber.Ctx) error {
		ret := map[string]any{
			"s": false,
			"m": "Method Not Allowed",
		}
		return ctx.Status(fiber.StatusMethodNotAllowed).JSON(ret)
	}, func(ctx *fiber.Ctx) bool {
		return ctx.Method() == fiber.MethodGet || ctx.Method() == fiber.MethodPost ||
			ctx.Method() == fiber.MethodPut || ctx.Method() == fiber.MethodDelete
	}))

	return fiberEngine
}

func SetupCompress(fiberEngine *fiber.App) {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add Compress Middleware")
	}
	fiberEngine.Use(compress.New(compress.Config{
		Next: func(c *fiber.Ctx) bool {
			compress := c.Get("compress", "1")
			return compress == "0"
		},
		Level: compress.LevelBestSpeed, // 1
	}))
}

func SetupCORS(fiberEngine *fiber.App) {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add CORS Middleware")
	}
	fiberEngine.Use(cors.New(cors.Config{
		// AllowOriginsFunc: func(origin string) bool {
		// 	fmt.Println(os.Getenv("ENVIRONMENT"))
		// 	return os.Getenv("ENVIRONMENT") == "development"
		// },
		AllowOrigins: os.Getenv("CORS_ALLOW_ORIGINS"),
		AllowHeaders: os.Getenv("CORS_ALLOW_HEADERS"),
		AllowMethods: os.Getenv("CORS_ALLOW_METHODS"),
	}))
}

func SetupLimiter(fiberEngine *fiber.App) {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Add Limiter Middleware")
	}
	fiberEngine.Use("/api", limiter.New(limiter.Config{
		// Next: func(c *fiber.Ctx) bool {
		// 	return c.IP() == "127.0.0.1" || c.IP() == "localhost"
		// },
		Max:        20,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			ret := map[string]any{
				"s": false,
				"m": "Too many request",
			}
			return c.Status(fiber.StatusTooManyRequests).JSON(ret)
		},
		// Storage: myCustomStorage{},
	}))
}

func SetupLogger(fiberEngine *fiber.App) {

	if log.IsTerminal(os.Stderr.Fd()) {
		log.DefaultLogger = log.Logger{
			TimeFormat: "15:04:05",
			Caller:     1,
			Writer: &log.ConsoleWriter{
				ColorOutput:    true,
				QuoteString:    true,
				EndWithMessage: true,
			},
		}
	}

	// create a new logger middleware with a custom writer
	fiberEngine.Use(fiberlog.New(&log.Logger{
		Context: log.NewContext(nil).Str("logger", "access").Value(),
		Writer: &log.FileWriter{
			Filename: "logs/access.log",
			MaxSize:  1024 * 1024 * 1024,
		},
	}, func(c *fiber.Ctx) bool {
		return string(c.Path()) == "/backdoor"
	}))

	// services.ConsoleLog("hello world")
	// services.ErrorLog("handle error")
	// services.DataLog("some data")
}
