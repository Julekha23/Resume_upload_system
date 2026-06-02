package setup

import (
	"First_Project/Setting"
	//"First_Project/upload"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// string connectionString =settings.Mysetting.Get_MongoURL()
var MongoClient *mongo.Client
func Connectbase(){
	cilentOption:=options.Client().ApplyURI(settings.Mysetting.Get_MongoURL())
	cilent,err:=mongo.Connect(context.TODO(),cilentOption)
	if err!=nil{
		panic(err)
	}
	MongoClient=cilent
	//upload.ResumeCollection = mongoClient.Database(settings.Mysetting.Get_mongo_DB()).Collection("resume")
	
}