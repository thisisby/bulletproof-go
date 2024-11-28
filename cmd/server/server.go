package main

import (
	"bulletproof-go/internal/core/config"
	"bulletproof-go/internal/core/context"
	"bulletproof-go/internal/core/di"
	"bulletproof-go/internal/core/middleware"
	"bulletproof-go/internal/core/module"
	"bulletproof-go/internal/core/router"
	"bulletproof-go/internal/modules/root"
	"bulletproof-go/internal/utils"
	"bulletproof-go/pkg/bytes"
	"bulletproof-go/pkg/cors"
	"log"
	"time"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()
	serverPort := cfg.Get("SERVER_PORT", "8080")
	env := cfg.Get("ENV", "development")
	addr := ":" + serverPort

	// CORS Options
	corsOptions := cors.SetupDefaultCorsOptions(
		[]string{"*"}, // Allow all origins
		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		[]string{"Content-Type", "Authorization"},           // Allowed headers
	)

	// Parser Options
	size := bytes.ConvertToBytes(int64(1024), bytes.Kilobytes)
	parserOptions := middleware.NewParserOptions(size)

	// redis client
	redisClient := utils.NewRedisClient("localhost:6379")

	// Initialize App with Middlewares
	App := router.NewRouter(
		router.WithCORS(*corsOptions),
		router.WithJSONParser(*parserOptions),
		router.WithCookieParser(),
		router.WithCsrf(),
		router.WithXss(),
		router.WithCaching(redisClient, 5*time.Minute, true),
		// router.WithInMemoryRateLimiter(4, 50, 1*time.Second, 5*time.Minute),
		router.WithRedisRateLimiter(redisClient, 100, 1*time.Second),
		// router.WithFileUpload("uploads", int64(1024), []string{"jpg", "jpeg", "png"}),
		router.WithProfiling(),
	)

	// Serve Static Files
	folderPath := utils.GetFolderPath("uploads")
	App.ServeStatic("/static/", folderPath)

	// Register dependencies
	dependencies := []interface{}{root.NewRootService, root.NewRootModule}
	di.RegisterDependencies(dependencies)

	// Root Module
	rootModule := root.NewRootModule(App)
	di.RegisterModules(App, []module.IModule{rootModule})

	// Example Route
	App.Get("/ping", func(ctx *context.Context) {
		ctx.Send("pong")
	})

	// Start the server
	log.Printf("Starting server on port %s in %s mode", serverPort, env)
	// LessGo.PProfiling()
	httpCfg := config.NewHttpConfig()
	if err := App.Listen(addr, httpCfg); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
