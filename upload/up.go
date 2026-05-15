package upload

import (
	settings "First_Project/Setting"
	"First_Project/model"
	setup "First_Project/set_up"
	"context"
	"os"
	"path"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ResumeCollection *mongo.Collection
func isValidName(name string) bool{
	for _,ch:=range name{
      if !unicode.IsLetter(ch) && !unicode.IsSpace(ch){
		return false
	  }
	}
	return true
}
func isValidPhone(phone string)bool{
	for _,ch:=range phone{
		if !unicode.IsDigit(ch){
			return false
		}
	}
	return true
}

func UploadResume(c *gin.Context) {
	fileHeader, err := c.FormFile("resume_file")
	if err != nil {
		c.JSON(200, gin.H{"error": "File not found"})
		return
	}
    if !strings.HasSuffix(strings.ToLower(fileHeader.Filename),".pdf"){
		c.JSON(400,gin.H{"eror":"only pdf files are allowed"})
		return
	}
	if fileHeader.Size > settings.Mysetting.Get_max_size()*1024*1024 {
		c.JSON(200, gin.H{"error": "file is too large"})
		return
	}

	candidateName := strings.TrimSpace(c.PostForm("candidate_name"))
	candidateEmail := strings.TrimSpace(c.PostForm("candidate_email"))
	phone := strings.TrimSpace(c.PostForm("phone"))
    if candidateName ==""{
		c.JSON(400,gin.H{"error":"candidate_name is required"})
		return 
	}
    if !isValidName(candidateName){
		c.JSON(400,gin.H{"error":"candidate_name must contain only letters"})
		return 
	}
	if !strings.Contains(candidateEmail,"@"){
		c.JSON(400,gin.H{"error":"candidate_email is not valid"})
	}
	if phone != "" && !isValidPhone(string(phone)){
		c.JSON(400,gin.H{"error":"phone must contain only numbers"})
	}
	uniqueId := uuid.New().String()
	savedFileName := uniqueId + ".pdf"
	storagePath := settings.Mysetting.Get_resume_path()
	fullPath := path.Join(storagePath, savedFileName)

	if err := os.MkdirAll("./storage/resumes", os.ModePerm); err != nil {
		c.JSON(200, gin.H{"error": "failed to create folder"})
		return
	}

	if err := c.SaveUploadedFile(fileHeader, fullPath); err != nil {
		c.JSON(200, gin.H{"error": "server error: cannot save file"})
		return
	}

	objID := primitive.NewObjectID()
	resume := &model.Resume{
		Id:               objID,
		ResumeID:         objID.Hex(),
		CandidateName:    candidateName,
		CandidateEmail:   candidateEmail,
		Phone:            phone,
		OriginalFileName: fileHeader.Filename,
		FilePath:         fullPath,
		FileSizeBytes:    fileHeader.Size,
	}

	_, err = ResumeCollection.InsertOne(context.TODO(), resume)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			c.JSON(200, gin.H{"error": "duplicate email or phone number"})
			return
		}
		c.JSON(200, gin.H{"error": "failed to save resume data"})
		return
	}

	c.JSON(200, gin.H{
		"message": "resume uploaded successfully",
		"resume":  resume,
	})
}
//indexing
func CreateResumeIndexes() {
	collection := setup.MongoClient.
		Database(settings.Mysetting.Get_mongo_DB()).
		Collection("resume")

	emailIndex := mongo.IndexModel{
		Keys:    bson.M{"candidate_email": 1},
		Options: options.Index().SetUnique(true),
	}
	phoneIndex := mongo.IndexModel{
		Keys:    bson.M{"phone": 1},
		Options: options.Index().SetUnique(true),
	}
	fileIndex:=mongo.IndexModel{
		Keys: bson.M{"original_file_name":1},
		Options: options.Index().SetUnique(true),
	}
	reviewerIndex:=mongo.IndexModel{
		Keys: bson.M{"reviewer_id":1},
		Options: options.Index().SetUnique(true).SetSparse(true),
	}

	_, err := collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		emailIndex,
		phoneIndex,
		fileIndex,
		reviewerIndex,
	})
	if err != nil {
		panic(err)
	}
}
