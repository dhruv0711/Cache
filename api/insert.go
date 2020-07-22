package api

import (
	"github.com/dhruv0711/Cache/cache"
	"github.com/dhruv0711/Cache/types"
	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	var todo types.Todo
	c.BindJSON(&todo)

	cache.InsertIntoCache(todo)

	c.JSON(201, "Inserted Successfully")
}
