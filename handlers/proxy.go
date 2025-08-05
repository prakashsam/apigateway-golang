package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
)

func ProxyHandler(target string) iris.Handler {
	return func(ctx iris.Context) {
		req := ctx.Request()
		fullURL := strings.TrimRight(target, "/") + ctx.Path()
		fmt.Println("Proxying request to:", fullURL)
		proxyReq, _ := http.NewRequest(req.Method, fullURL, req.Body)
		proxyReq.Header = req.Header

		traceHeader := ctx.GetHeader("X-Cloud-Trace-Context")
		proxyReq.Header.Set("X-Cloud-Trace-Context", traceHeader)

		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			ctx.StatusCode(iris.StatusBadGateway)
			ctx.WriteString("Proxy Error")
			return
		}
		body, _ := io.ReadAll(resp.Body)
		ctx.StatusCode(resp.StatusCode)
		ctx.Write(body)
	}
}
