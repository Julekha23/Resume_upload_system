package controller

import (
	settings "First_Project/Setting"
	"First_Project/model"
	setup "First_Project/set_up"
	"First_Project/upload"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ReviewerCollection *mongo.Collection
func InitCollection(){
	ReviewerCollection=setup.MongoClient.Database(settings.Mysetting.Get_mongo_DB()).Collection(settings.Mysetting.Get_Reviewer_DB())
}
func AddReviewer(c *gin.Context){
	var body struct {
		ReviewerID string `json:"reviewer_id"`
		Name string `json:"name"`
		Email string `json:"email"`
	}
	if err:=c.ShouldBindJSON(&body);err!=nil{
		c.JSON(400,gin.H{"error":"invalid request body"})
		return 
	}
	if body.ReviewerID ==""||body.Name==""||body.Email==""{
		c.JSON(400,gin.H{"error":"reviewer_id,name,email are required"})
		return 
	}
	reviewer:=model.Reviewer{
		Id: primitive.NewObjectID(),
		ReviewerID: body.ReviewerID,
		Name:body.Name,
		Email:body.Email,
		Status: model.ReviewerActive,
	}
	ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	_,err:=ReviewerCollection.InsertOne(ctx,reviewer)
	if err!=nil{
		if mongo.IsDuplicateKeyError(err){
			c.JSON(409,gin.H{"error":"reviewer_id already exists"})
			return 
		}
		c.JSON(500,gin.H{"error":"failed to fetch reviewers"})
	    return
	}
	c.JSON(200,gin.H{
		"message":"reviewer added successfully",
		"reviewer":reviewer,
	})
}
func GetAllReviewers(c *gin.Context){
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	curser,err :=ReviewerCollection.Find(ctx,bson.M{})
	if err!=nil{
		c.JSON(500,gin.H{"error":"failed to fetch reviewers"})
		return 

	}
	var reviewers []model.Reviewer
	if err=curser.All(ctx,&reviewers);err!=nil{
		c.JSON(500,gin.H{"error":"failed to decode reviewers"})
		return 
	}
	c.JSON(200,gin.H{"data":reviewers})
}
func UpdateStatusAndFreeReviewer(c *gin.Context){
	objID,err :=primitive.ObjectIDFromHex(c.Param("id"))
	if err!=nil{
		c.JSON(400,gin.H{"error":"Invalid resume id"})
		return 
	}
	var body struct {
		Status string `json:"status"`
	}
	if err :=c.ShouldBindJSON(&body);err!=nil{
		c.JSON(400,gin.H{"error":"invalid request body"})
		return 
	}
	validStatus :=map[string]bool{
		"pending":   true,
		"reviewing": true,
		"done":      true,
		"rejected":  true,
	}
	if !validStatus[body.Status]{
		c.JSON(400,gin.H{"error":"invalid status.allowed:pending,reviewing,done,rejected"})
		return 
	}
	ctx,cancel :=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	var resume model.Resume
	err =upload.ResumeCollection.FindOne(ctx,bson.M{"_id":objID}).Decode(&resume)
	if err!=nil{
		c.JSON(404,gin.H{"error":"resume not found"})
		return 
	}
	result,err:=upload.ResumeCollection.UpdateOne(
		ctx,
		bson.M{"_id":objID},
		bson.M{"$set":bson.M{"status":body.Status},
	           "$unset":bson.M{"reviewer_id":""},                                 },
	)
	if err!=nil{
		c.JSON(500,gin.H{"error":"failed to update status"})
		return 
	}
	if result.MatchedCount==0{
		c.JSON(404,gin.H{"error":"resume not found"})
		return 
	}
	if body.Status=="done" || body.Status=="rejected" {
		if resume.ReviewerID != ""{
			_,err:=ReviewerCollection.UpdateOne(
				ctx,
				bson.M{"reviewer_id":resume.ReviewerID},
				bson.M{"$set":bson.M{"status":model.ReviewerActive}},
			)
			_,err=upload.ResumeCollection.UpdateOne(
				ctx,
				bson.M{"resume_id":resume.ResumeID},
				bson.M{"$unset":bson.M{"reviewer_id":""}},
			)
		if err!=nil{
			c.JSON(200,gin.H{
				"message":"status updated but could not free review",
				"status":body.Status,
			})
			return
		}
		}
	}
	c.JSON(200,gin.H{
		"message":"status updated successfully",
		"status":body.Status,
	})
}
func AssignReviewer(c *gin.Context){
	ObjId,err:=primitive.ObjectIDFromHex(c.Param("id"))
	if err!=nil{
		c.JSON(400,gin.H{"error":"Invalid resume id"})
		return 
	}
	var body struct{
		Reviewer_Id string `json:"reviewer_id"`
	}
	if err:=c.ShouldBindJSON(&body);
	err!=nil{
		c.JSON(400,gin.H{"error" :"Invalid  request body",})
		return 
	}
	if body.Reviewer_Id ==""{
		c.JSON(400,gin.H{"error":"Reviewer id is required..."})
		return 
	}
    ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	var rev model.Reviewer
	err=ReviewerCollection.FindOne(ctx,bson.M{
		"reviewer_id":body.Reviewer_Id,
		"status":model.ReviewerActive,
	}).Decode(&rev)
    if err!=nil{
		c.JSON(400,gin.H{"error":"reviewer  not found else reviewer is busy"})
		return 
	}
	_,err=ReviewerCollection.UpdateOne(ctx,bson.M{"reviewer_id":body.Reviewer_Id},bson.M{"$set":bson.M{"status":model.ReviewerInactive},})
	if err!=nil{
		c.JSON(400,gin.H{"error":"Failed to update reviewer status"})
		return 
	}
	result,err:=upload.ResumeCollection.UpdateOne(ctx,bson.M{"_id":ObjId},bson.M{"$set":bson.M{"reviewer_id":body.Reviewer_Id,"status":"reviewing",}},)
	if err!=nil{
		c.JSON(400,gin.H{"error":"Failed to assign reviewer."})
		return 
	}
	if result.MatchedCount==0{
		c.JSON(500,gin.H{"error":"failed to assign reviewer"})
		return
	}
	c.JSON(200,gin.H{
		"message": "reviewer assigned successfully",
		"reviewer_id":body.Reviewer_Id,
		"status":"reviewing",
	})
}