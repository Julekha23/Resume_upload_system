package controller

import (
	"First_Project/model"
	"First_Project/upload"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/twitchyliquid64/golang-asm/obj"
	//"github.com/twitchyliquid64/golang-asm/obj"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

func Start(c *gin.Context) {
	fmt.Println("Resume uploaded system")
	c.JSON(200, gin.H{"message": "welcome..."})
}

func Display(c *gin.Context) {
	cursor, err := upload.ResumeCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(200, gin.H{"error": "failed to fetch data"})
		return
	}

	var resumes []model.Resume
	if err = cursor.All(context.TODO(), &resumes); err != nil {
		c.JSON(200, gin.H{"error": "failed to decode data"})
		return
	}

	c.JSON(200, gin.H{"data": resumes})
}
func GetByID(c *gin.Context){
	idParam :=c.Param("id")
	objID,err :=primitive.ObjectIDFromHex(idParam)
	if err!=nil{
		c.JSON(400,gin.H{"error":"invalid resume id"})
		return 
	}
	var resume model.Resume
	err =upload.ResumeCollection.FindOne(context.TODO(),bson.M{"_id":objID}).Decode(&resume)
	if err!=nil{
		c.JSON(404,gin.H{"error":"resume not found"})
		return 
	}
	c.JSON(200,gin.H{"data":resume})
}
func DeleteByID(c *gin.Context){
	idParam := c.Param("id")
	objId,err:=primitive.ObjectIDFromHex(idParam)
	if err!=nil{
		c.JSON(400,gin.H{"error":"invalid resume id"})
		return 
	}
	var resume model.Resume
	err =upload.ResumeCollection.FindOneAndDelete(context.TODO(),bson.M{"_id":objId}).Decode(&resume)
	if err!=nil{
		c.JSON(500,gin.H{"error":"resume not found and failed to delete"})
		return 
	}
	if err :=os.Remove(resume.FilePath); err !=nil{
		fmt.Println("Warning: could not delete file:",resume.FilePath)
	}
	c.JSON(200,gin.H{"message":"resume deleted successfully"})
}
func DownloadFile(c *gin.Context){
	idParam := c.Param("id")
	objID,err:=primitive.ObjectIDFromHex(idParam)
	if err !=nil{
		c.JSON(400,gin.H{"error":"invalid resume id"})
		return 
	}
	var resume model.Resume
	err =upload.ResumeCollection.FindOne(context.TODO(),bson.M{"_id":objID}).Decode(&resume)
	if err!=nil{
		c.JSON(404,gin.H{"error":"resume not found"})
		return 
	}
	if _,err:=os.Stat(resume.FilePath);os.IsNotExist(err){
		c.JSON(404,gin.H{"error":"file not found on server"})
		return 
	}
	c.Header("Content-Dispoaition","attachment;filename="+resume.OriginalFileName)
	c.Header("Content-Type","application/pdf")
	c.File(resume.FilePath)
}
// func UpdateStatus(c *gin.Context){
// 	objID,err:=primitive.ObjectIDFromHex(c.Param("id"))
// 	if err!=nil{
// 		c.JSON(400,gin.H{"error":"invalid resume id"})
// 		return 
// 	}
// 	var body struct{
// 		Status string `json:"status"`
// 	}
// 	if err :=c.ShouldBindJSON(&body);err!=nil{
// 		c.JSON(400,gin.H{"error":"invalid request body"})
// 		return 
// 	}
// 	validStatuses :=map[string] bool{
// 		"pending":true,
// 		"reviewing":true,
// 		"done":true,
// 		"rejected":true,
// 	}
// 	if !validStatuses[body.Status] {
// 		c.JSON(400,gin.H{
// 			"error":"Invalid status.allowed values:pending,reviewing,done,rejected",
// 		})
// 		return 
// 	}
// 	result,err:= upload.ResumeCollection.UpdateOne(
// 		context.TODO(),
// 		bson.M{"_id":objID},
// 		bson.M{"$set":bson.M{"status":body.Status}},
// 	)
// 	if err!=nil{
// 		c.JSON(500,gin.H{"Error":"failed to update status"})
// 	}
// 	if result.MatchedCount==0{
// 		c.JSON(404,gin.H{"error":"resume not found"})
// 		return 
// 	}
// 	c.JSON(200,gin.H{
// 		"message":"status updated successfully",
// 		"status":body.Status,
// 	})
// }
// // func AssignReviewer(c *gin.Context){
// // 	objID,err:=primitive.ObjectIDFromHex(c.Param("id"))
// // 	if err!=nil{
// // 		c.JSON(400,gin.H{"error":"invalid resume id"})
// // 		return 
// // 	}
// // 	var body struct{
// // 		ReviewerID string `json:"reviewer_id"`
// // 	}
// // 	if err :=c.ShouldBindJSON(&body); err!=nil{
// // 		c.JSON(400,gin.H{"error":"Invalid request body"})
// // 		return 
// // 	}
// // 	if body.ReviewerID==""{
// // 		c.JSON(400,gin.H{"error":"Invalid request body"})
// // 		return 
// // 	}
// // 	result,err :=upload.ResumeCollection.UpdateOne(
// // 		context.TODO(),
// // 		bson.M{"_id":objID},
// // 		bson.M{"$set":bson.M{
// // 			"reviewer_id":body.ReviewerID,
// // 			"status":"reviewing",
// // 		}},
// // 	)
// // 	if err!=nil{
// // 		if mongo.IsDuplicateKeyError(err){
// // 			c.JSON(409,gin.H{"error":"this reviewer is already assigned to another resume"})
// // 			return 
// // 		}
// // 		c.JSON(500,gin.H{"error":"failed to assign reviewer"})
// // 		return 
// // 	}
// // 	if result.MatchedCount == 0{
// // 		c.JSON(404,gin.H{"error":"resume not found"})
// // 		return 
// // 	}
// // 	c.JSON(200,gin.H{
// // 		"message":"reviewer assigned successfully",
// // 		"reviewer_id":body.ReviewerID,
// // 		"status":"reviewing",
// // 	})
// // }