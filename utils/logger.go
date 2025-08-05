package utils

import (
    "context"
    "log"

    "cloud.google.com/go/logging"
)

func SetupLogger(projectID string) *logging.Logger {
    client, err := logging.NewClient(context.Background(), projectID)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    return client.Logger("api-gateway")
}