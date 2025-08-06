package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Teacher struct {
	TId    int
	Name   string
	Age    int
	Gender string
}

type Student struct {
	SId  int
	Name string
	Age  int
}

func main() {
	r := gin.Default()
	t := Teacher{}
	s := Student{}
	r.POST("/test", func(c *gin.Context) {
		if errA := c.ShouldBind(&t); errA != nil {
			c.String(http.StatusOK, "the body should be teacher")
		} else if errB := c.ShouldBind(&s); errB != nil {
			c.String(http.StatusOK, "the body should be student")
		}
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
