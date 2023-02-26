package controllers

import (
	"context"
	"fmt"
	"gdsc/configs"
	"gdsc/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)
var blogsCollection = configs.GetCollection(configs.DB, "blog")

func BlogPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var blogPayload models.Blog
		defer cancel()
		c.BindJSON(&blogPayload)


		newPoint := models.Blog{
			Title: blogPayload.Title,
			Content: blogPayload.Content,
		}

		result, err := blogsCollection.InsertOne(ctx, newPoint)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		fmt.Println(result)

		finalView := models.Blog{
			Title:  blogPayload.Title,
			Content: blogPayload.Content,
		}

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusCreated,
			"message": "success",
			"data":    finalView,
		})

	}
}