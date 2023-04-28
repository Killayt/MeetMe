package server

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", indexHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.LoadHTMLFiles("web/frontend")
}

func indexHandler(c *gin.Context) {
	c.File("web/frontend/index.html")
}

// func indexHandler(c *gin.Context) {
//   // Send a 200 status code and the message "Hello, World!" in the response body
//   c.String(http.StatusOK, "Hello, World!")
// }
