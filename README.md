# 🚀 Gin + Svelte (TS & Vite)

A starter template to start a new fullstack web application with Gin web framework as a backend and Svelte as a frontend. Prepared as a monorepo.

## Prerequisites

- Machine with MacOS or Linux
- Go lang installed
- Docker installed
- Docker compose installed

## Steps to reproduce

### Step 1: Initiate the go project

```bash
go mod init
```

### Step 2: Initiate the server

```bash
touch server.go
```
Then add this line:

```go
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ack": time.Now(),
		})
	})

	r.Run(":8080")
}
```

Then run:

```bash
go mod tidy
```

Give it a try:

```bash
go run server.go
# then open http://localhost:8080 in the browser
```

### Step 3: Add Client

```bash
mkdir client
cd client
npm init vite
```

Let Gin serve the client's static files:

```go
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

	...

	r.GET("/api/random", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"random": time.Now().UnixNano(),
		})
	})

	r.Run(":8080")
}
```
