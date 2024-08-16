package main

import (
	"fmt"
	"net/http"
	"os"
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

	// Routes
	r.GET("/", homeHandler)
	r.GET("/about", aboutHandler)
	r.GET("/contact", contactHandler)
	r.GET("/data", dataHandler)
	r.GET("/external", externalAPIHandler)
	r.GET("/complex-query", complexQueryHandler)
	r.GET("/cache", cacheHandler)
	r.GET("/file", fileIOHandler)
	r.GET("/unauthorized", unauthorizedHandler)
	r.GET("/forbidden", forbiddenHandler)
	r.GET("/not-found", notFoundHandler)
	r.GET("/internal-error", internalErrorHandler)

	r.Run(":9999")
}

func homeHandler(c *gin.Context) {
	c.String(200, "Welcome to ob-test!")
}

func aboutHandler(c *gin.Context) {
	c.String(200, "Mock application.\n\nHere are the endpoint names:\n"+
		"1. `/`\n"+
		"2. `/about`\n"+
		"3. `/contact`\n"+
		"4. `/data`\n"+
		"5. `/external`\n"+
		"6. `/complex-query`\n"+
		"7. `/cache`\n"+
		"8. `/file`\n"+
		"9. `/unauthorized`\n"+
		"10. `/forbidden`\n"+
		"11. `/not-found`\n"+
		"12. `/internal-error`.")
}

func contactHandler(c *gin.Context) {
	c.String(200, "Contact us at: initializ.ai")
}

func dataHandler(c *gin.Context) {
	// Simulating a database call
	time.Sleep(500 * time.Millisecond) // Simulate latency
	fakeDBResponse := queryDatabase()

	c.JSON(200, gin.H{
		"data": fakeDBResponse,
	})
}

func queryDatabase() string {
	// Simulate a database query
	time.Sleep(200 * time.Millisecond) // Simulate latency
	return "Fake DB response: {'id': 1, 'name': 'Test Item'}"
}

func externalAPIHandler(c *gin.Context) {
	// Simulating an external API call
	time.Sleep(1 * time.Second)
	resp, err := http.Get("https://console.dev.initializ.ai/login/")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to call external API",
		})
		return
	}
	defer resp.Body.Close()

	c.JSON(resp.StatusCode, gin.H{
		"external_data": fmt.Sprintf("Received data from external API with status code %d", resp.StatusCode),
	})
}

func complexQueryHandler(c *gin.Context) {
	time.Sleep(300 * time.Millisecond) // Simulate initial query latency
	result1 := queryDatabase()

	time.Sleep(300 * time.Millisecond) // Simulate processing latency
	result2 := anotherDatabaseQuery()

	c.JSON(200, gin.H{
		"complex_data": fmt.Sprintf("Combined Results: %s, %s", result1, result2),
	})
}

func anotherDatabaseQuery() string {
	time.Sleep(200 * time.Millisecond) // Simulate latency
	return "Another DB response: {'id': 2, 'name': 'Another Item'}"
}

func cacheHandler(c *gin.Context) {
	time.Sleep(100 * time.Millisecond)
	cacheData := "Cached data: {'key': 'cached_value'}"

	c.JSON(200, gin.H{
		"cache_data": cacheData,
	})
}

func fileIOHandler(c *gin.Context) {
	filePath := "./sample.txt"
	err := os.WriteFile(filePath, []byte("Hello, World!"), 0644)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to write to file",
		})
		return
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to read from file",
		})
		return
	}
	defer os.Remove(filePath)
	c.JSON(200, gin.H{
		"file_content": string(data),
	})
}

func unauthorizedHandler(c *gin.Context) {
	// Simulating an unauthorized request
	c.JSON(401, gin.H{
		"error": "Unauthorized access",
	})
}

func forbiddenHandler(c *gin.Context) {
	// Simulating a forbidden request
	c.JSON(403, gin.H{
		"error": "Forbidden access",
	})
}

func notFoundHandler(c *gin.Context) {
	// Simulating a not found error
	c.JSON(404, gin.H{
		"error": "Resource not found",
	})
}

func internalErrorHandler(c *gin.Context) {
	// Simulating an internal server error
	c.JSON(500, gin.H{
		"error": "Internal server error",
	})
}
