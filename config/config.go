package config

import "os"

type Config struct {
    OrderServiceURL  string
    PaymentServiceURL string
    AuthServiceURL   string
    JWTSecret        string
    ProjectID        string
}

func Load() Config {
    return Config{
        OrderServiceURL:  os.Getenv("ORDER_SERVICE_URL"),
        PaymentServiceURL: os.Getenv("PAYMENT_SERVICE_URL"),
        AuthServiceURL:   os.Getenv("AUTH_SERVICE_URL"),
        JWTSecret:         os.Getenv("JWT_SECRET"),
        ProjectID:         os.Getenv("GCP_PROJECT_ID"),
    }
}