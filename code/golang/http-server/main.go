package main

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	env1 "github.com/codingconcepts/env"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// EnvConfig
type EnvConfig struct {
	Tag           string `env:"TAG" required:"true"`
	Port          string `env:"PORT" required:"true"`
	APIProxy      string `env:"API_PROXY" required:"true"`
	APIPathPrefix string `env:"API_PATH_PREFIX" required:"true"`
	LogPath       string `env:"LOG_PATH" required:"true"`
}

var (
	//go:embed resources/*
	//go:embed resources/_next/static/chunks/pages/_app*
	//go:embed resources/_next/static/*/_*
	staticFs embed.FS

	// envConfig
	envConfig *EnvConfig

	// logger
	logger *logrus.Logger
)

func init() {
	envConfig = &EnvConfig{}
	if err := env1.Set(envConfig); err != nil {
		panic(fmt.Errorf("parse env file error: %s", err.Error()))
	}
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logPath := fmt.Sprintf("%s/%s.api.log", envConfig.LogPath, "%Y-%m-%d_%H")
	writer, err := rotatelogs.New(logPath,
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(fmt.Sprintf("new rotatelogs error: %s", err.Error()))
	}
	logger.SetOutput(writer)
}

func main() {

	router := gin.Default()
	router.Use(apiProxy(envConfig.APIProxy), cookie)
	sub, err := fs.Sub(staticFs, "resources")
	if err != nil {
		panic("resources not found")
	}
	router.StaticFS("/", http.FS(sub))

	router.Run(":" + envConfig.Port)
}

func cookie(ctx *gin.Context) {
	if strings.HasSuffix(ctx.Request.URL.Path, ".html") {
		ctx.SetCookie("tag", envConfig.Tag, 3600*24*365, "/", ctx.Request.Host, false, false)
	}
}

func apiProxy(proxy string) gin.HandlerFunc {
	apiProxyURL, err := url.Parse(proxy)
	if err != nil {
		panic(fmt.Sprintf("parse api proxy error: %s", err.Error()))
	}
	return func(ctx *gin.Context) {

		if !strings.HasPrefix(ctx.Request.URL.Path, envConfig.APIPathPrefix) {
			ctx.Next()
			return
		}
		logAPI(ctx)

		urlObject := &url.URL{
			Scheme: apiProxyURL.Scheme,
			Host:   apiProxyURL.Host,
		}
		apiProxyPrefix := strings.Trim(apiProxyURL.Path, "/")

		proxy := httputil.NewSingleHostReverseProxy(urlObject)

		realPath := strings.Replace(ctx.Request.URL.Path, envConfig.APIPathPrefix, "", -1)
		if len(apiProxyPrefix) > 0 {
			realPath = "/" + apiProxyPrefix + realPath
		}
		ctx.Request.URL.Path = realPath
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
		ctx.Abort()
	}
}

func logAPI(ctx *gin.Context) {
	data := map[string]interface{}{
		"url": ctx.Request.URL.String(),
	}
	headers := make(map[string]string)
	for k, v := range ctx.Request.Header {
		headers[k] = strings.Join(v, ",")
	}
	data["header"] = headers

	if raw, err := ctx.GetRawData(); err == nil {
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(raw))
		data["body"] = string(raw)
	}

	logger.WithFields(logrus.Fields(data)).Info("api")

}
