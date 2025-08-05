package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"api-gateway/middleware"
	"api-gateway/utils"

	"github.com/kataras/iris/v12"
)

func main() {
	godotenv.Load()
	cfg := config.Load()
	app := iris.New()
	app.Use(middleware.LoggingMiddleware(cfg.ProjectID))
	app.Use(middleware.TraceMiddleware())
	fmt.Println("Order service URL:", cfg.OrderServiceURL)
	jwtSecret, err := utils.GetSecret("JWT_SECRET", cfg.ProjectID)
	fmt.Println("jwtSecret:", jwtSecret)
	if err != nil {
		log.Fatalf("Failed to get JWT secret: %v", err)
	}

	app.Any("/auth/{p:path}", handlers.ProxyHandler(cfg.AuthServiceURL))

	app.Use(middleware.JWTAuth(jwtSecret))
	app.Any("/orders", handlers.ProxyHandler(cfg.OrderServiceURL))
	app.Any("/orders/{p:path}", handlers.ProxyHandler(cfg.OrderServiceURL))
	app.Any("/payment/{p:path}", handlers.ProxyHandler(cfg.PaymentServiceURL))

	app.Listen(":8080")
}
