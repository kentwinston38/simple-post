package controller

import (
	"myapp/model"
	"myapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var (
		input model.NewPost
	)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.PostReturn{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	post, err := service.PostCreate(c.Request.Context(), input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.PostReturn{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &model.PostReturn{
		Success: true,
		Message: "Success",
		Data:    post,
	})
}

func PostGetAll(c *gin.Context) {
	posts, err := service.PostGetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.PostBatchReturn{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &model.PostBatchReturn{
		Success: true,
		Message: "Success",
		Data:    posts,
	})
}
