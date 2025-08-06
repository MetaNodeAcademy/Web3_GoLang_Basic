package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	t := Person{}
	gin.ForceConsoleColor()
	r.GET("/test/:name/:id", func(c *gin.Context) {
		if errA := c.ShouldBindUri(&t); errA != nil {
			c.JSON(400, gin.H{"msg": errA.Error()})
		}
		c.JSON(http.StatusOK, t)
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
