package upload

import (
	settings "First_Project/Setting"
	"First_Project/model"
	setup "First_Project/set_up"
	"context"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var ResumeCollection *mongo.Collection

func UploadResume(c *gin.Context) {
	fileHeader, err := c.FormFile("resume_file")
	if err != nil {
		c.JSON(200, gin.H{
			"error": "File not found",
		})
		return
	}
	if fileHeader.Size > settings.Mysetting.Get_max_size()*1024*1024 {
		c.JSON(200, gin.H{
			"error": "file is too large",
		})
		return
	}
	candidateName := c.PostForm("candidate_name")
	candidateEmail := c.PostForm("candidate_email")
	phone := c.PostForm("phone")
	uniqueId:= uuid.New().String()
	savedFileName:= uniqueId+".pdf"
	storagePath:= settings.Mysetting.Get_resume_path()
	fullPath := path.Join(storagePath, savedFileName)
	if err:=os.MkdirAll("./storage/resumes", os.ModePerm);err!=nil{
		c.JSON(200, gin.H{
		"error": "failed to create folder",
	    })
	    return
	}
	if err := c.SaveUploadedFile(fileHeader, fullPath); err != nil {
	    c.JSON(200, gin.H{
		"error": "server error: cannot save file",
	    })
	    return
	}
	objID := primitive.NewObjectID()
	resume :=&model.Resume{
		Id: objID,
		ResumeID:objID.Hex(),
		CandidateName:candidateName,
		CandidateEmail: candidateEmail,
		Phone: phone,
		OriginalFileName:fileHeader.Filename,
		FilePath :fullPath,
		FileSizeBytes:fileHeader.Size,
	}
	// ResumeCollection:=setup.MongoClient.Database(settings.Mysetting.Get_mongo_DB()).Collection("resume")
	// _, err = ResumeCollection.InsertOne(context.TODO(),resume)
    // if err != nil {
	// c.JSON(200, gin.H{
	// 	"error": "failed to save resume data",
	// })
	// return
    // }
	_, err = ResumeCollection.InsertOne(context.TODO(), resume)
    if err != nil {
	if mongo.IsDuplicateKeyError(err) {
		c.JSON(200, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"error": "failed to save resume data",
	})
	return
}
    c.JSON(200, gin.H{
	"message": " resume uploaded successfully ",
	"resume":  resume,
    })
}
func CreateResumeIndexes() {

	collection := setup.MongoClient.
		Database(settings.Mysetting.Get_mongo_DB()).
		Collection("resume")

	emailIndex := mongo.IndexModel{
		Keys: bson.M{
			"candidate_email": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	nameIndex:=mongo.IndexModel{
		Keys:bson.M{
			"candidate_name":1,
		},
		Options:options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().
		CreateMany(context.TODO(), []mongo.IndexModel{
			emailIndex,
			nameIndex,
		})

	if err != nil {
		panic(err)
	}
}