package middleware

import "github.com/kataras/iris/v12"

func TraceMiddleware() iris.Handler {
    return func(ctx iris.Context) {
        trace := ctx.GetHeader("X-Cloud-Trace-Context")
        if trace != "" {
            ctx.ResponseWriter().Header().Set("X-Cloud-Trace-Context", trace)
        }
        ctx.Next()
    }
}