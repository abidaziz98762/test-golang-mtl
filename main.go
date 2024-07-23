package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Log format: timestamp, method, path, status, latency, IP
		logger.Infof("[%s] \"%s %s\" %d \"%s\" \"%s\"",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
		)
		return ""
	}), gin.Recovery())

	r.GET("/", homeHandler)
	r.GET("/about", aboutHandler)
	r.GET("/contact", contactHandler)

	r.Run(":9999")
}

func homeHandler(c *gin.Context) {
	c.String(200, "Welcome to the home page!")
}

func aboutHandler(c *gin.Context) {
	c.String(200, "This is the about page.")
}

func contactHandler(c *gin.Context) {
	c.String(200, "Contact us at: initializ.ai")
}
