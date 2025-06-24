package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"him/fiber-api/core"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

//go:embed public/** assets/**
var embedFile embed.FS

//go:embed assets/certs/ssl.cert
var certPEM []byte

//go:embed assets/certs/ssl.key
var keyPEM []byte

//go:embed assets/db/app.db
var dbFile []byte

func writeTempDB() (string, error) {
	tmpFile, err := os.CreateTemp("", "app-*.db")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	if _, err := tmpFile.Write(dbFile); err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}

func main() {
	dbPath, err := writeTempDB()
	if err != nil {
		log.Fatalf("write temp db: %v", err)
	}
	defer os.Remove(dbPath) // optional: remove after run

	app := core.SetupApp(embedFile, dbPath)
	app.Get("/metrics", monitor.New(monitor.Config{Title: os.Getenv("APP_NAME")}))

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedFile),
		PathPrefix: "public", // important for folder structure
		Browse:     false,
		Index:      "index.html",
	}))
	// app.Static("/", "./public")
	// app.Static("/_nuxt", "app/.output/public/_nuxt")
	// app.Static("/themes", "app/.output/public/_nuxt")
	// app.Get("/app/*", func(ctx *fiber.Ctx) error {
	// 	return ctx.SendFile("app/.output/public/index.html")
	// })
	// Optional: Redirect root to index.html
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendFile("./public/index.html")
	// })
	// SPA fallback (route like /dashboard, /settings)
	app.Use("*", func(c *fiber.Ctx) error {
		data, err := embedFile.Open("public/index.html")
		if err != nil {
			return c.Status(404).SendString("404 Not Found")
		}
		c.Set("Content-Type", "text/html") // <- penting!
		return c.SendStream(data)
	})
	// data, err := publicFile.ReadFile("public/index.html")
	// app.Get("*", func(c *fiber.Ctx) error {
	// 	return c.se (http.FS(publicFile))
	// })

	// app.Listen(":3000")

	// Create tls certificate
	// cer, err := tls.LoadX509KeyPair("assets/certs/ssl.cert", "assets/certs/ssl.key")
	cer, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	// Create custom listener
	ln, err := tls.Listen("tcp", ":"+os.Getenv("APP_PORT"), config)
	if err != nil {
		panic(err)
	}
	// app.Listener(ln)

	// Listen from a different goroutine
	go func() {
		// if err := app.Listen(":3000"); err != nil {
		if err := app.Listener(ln); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = ln.Close()
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
