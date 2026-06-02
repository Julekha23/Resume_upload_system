package controller

import (
	"First_Project/model"
	// setup "First_Project/set_up"
	"First_Project/upload"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Start(c *gin.Context){
	fmt.Println("Resume uploaded system")
	c.JSON(200,gin.H{"message":"welcome..."})
}
func Display(c *gin.Context){
	cursor, err := upload.ResumeCollection.Find(context.TODO(),bson.M{},)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "failed to fetch data",
		})
		return
	}

	var resumes []model.Resume

	if err = cursor.All(context.TODO(), &resumes); err != nil {
		c.JSON(200, gin.H{
			"error": "failed to decode data",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": resumes,
	})
}