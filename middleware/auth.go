package middleware

import (
    "strings"

    "github.com/golang-jwt/jwt/v4"
    "github.com/kataras/iris/v12"
)

func JWTAuth(secret string) iris.Handler {
    return func(ctx iris.Context) {
        authHeader := ctx.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            ctx.StatusCode(401)
            return
        }
        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })
        if err != nil || !token.Valid {
            ctx.StatusCode(401)
            return
        }
        ctx.Next()
    }
}