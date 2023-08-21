package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./client/**/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Static("/assets", "./client/dist/assets")
	r.StaticFile("/vite.svg", "./client/dist/vite.svg")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ack": time.Now(),
		})
	})

	r.GET("/api/random", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"random": time.Now().UnixNano(),
		})
	})

	r.Run(":8080")
}
