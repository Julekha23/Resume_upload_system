package main

import (
	settings "First_Project/Setting"
	"First_Project/controller"
	setup "First_Project/set_up"
	"First_Project/upload"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	settings.Generate()
	setup.Connectbase()
	fmt.Println("Resume Upload System Started")
	upload.ResumeCollection = setup.MongoClient.
		Database(settings.Mysetting.Get_mongo_DB()).
		Collection("resume")
	controller.InitCollection()	
	upload.CreateResumeIndexes()
	// Set up routes
	r := gin.Default()
	r.GET("/", controller.Start)
	r.PUT("/resume", upload.UploadResume)
	r.GET("/resume", controller.Display)
	r.GET("/resume/:id",controller.GetByID)
	r.DELETE("/resume/:id",controller.DeleteByID)
	r.GET("/resume/:id/file",controller.DownloadFile)
	// r.PATCH("/resume/:id/status",controller.UpdateStatus)
	// r.PATCH("/resume/:id/assign",controller.AssignReviewer)
	r.PATCH("/resume/:id/status",controller.UpdateStatusAndFreeReviewer)
	r.PATCH("/resume/:id/assign",controller.AssignReviewer)
	r.POST("/reviewer",controller.AddReviewer)
	r.GET("/reviewer",controller.GetAllReviewers)
	// Start server
	r.Run(fmt.Sprintf(":%d", settings.Mysetting.Get_Server()))
}
