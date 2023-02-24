package controllers

import (
	"go_crud/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context){
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.CreatedAt = time.Now()
	models.DB.NewInsert().Model(&comment).Exec(c)

	c.JSON(http.StatusOK, gin.H{"msg": "Successfully commented!"})
}