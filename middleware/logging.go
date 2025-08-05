package middleware

import (
    "context"
    "log"

    "cloud.google.com/go/logging"
    "github.com/kataras/iris/v12"
)

func LoggingMiddleware(projectID string) iris.Handler {
    client, err := logging.NewClient(context.Background(), projectID)
    if err != nil {
        log.Fatalf("Failed to create logging client: %v", err)
    }
    logger := client.Logger("api-gateway")

    return func(ctx iris.Context) {
        logger.Log(logging.Entry{
            Severity: logging.Info,
            Payload: map[string]interface{}{
                "path":       ctx.Path(),
                "method":     ctx.Method(),
                "statusCode": ctx.GetStatusCode(),
                "trace":      ctx.GetHeader("X-Cloud-Trace-Context"),
            },
        })
        ctx.Next()
    }
}