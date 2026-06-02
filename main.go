package main

import (
	"First_Project/Setting"
	"First_Project/controller"
	setup "First_Project/set_up"
	"First_Project/upload"
	//"context"
	"fmt"
	//"log"
	//"time"

	"github.com/gin-gonic/gin"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)
func init(){
    settings.Generate()
	setup.Connectbase()
	fmt.Println("Resume Upload")
	upload.ResumeCollection = setup.MongoClient.
	Database(settings.Mysetting.Get_mongo_DB()).
	Collection("resume")
	upload.CreateResumeIndexes()
	r:=gin.Default()
	r.GET("/",controller.Start)
	r.PUT("/resume",upload.UploadResume)
	r.GET("/resume",controller.Display)
	r.Run(fmt.Sprintf(":%d",settings.Mysetting.Get_Server()))
}
func main() {
	fmt.Println("Hello")
}