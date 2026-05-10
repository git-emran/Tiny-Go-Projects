package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Name  string `json: "name"`
	Email string `json: "email" binding: "required, email"`
}

func router() *gin.Engine {
	r := gin.Default()
	userRoute := r.Group("/user")
	userRoute.GET("/hello/:name", func(c *gin.Context) {
		user := c.Param("name")
		response := fmt.Sprintf("Hello, %s", user)
		c.String(http.StatusOK, response)
	})

	userRoute.POST("/post", func(c *gin.Context) {
		body := Message{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
	})

	return r
}

func main() {
	router().Run()
}
