package setup

import (
<<<<<<< HEAD
	"First_Project/Setting"
	//"First_Project/upload"
	"context"
=======
	settings "First_Project/Setting"
	"context"
	"fmt"
	"time"
>>>>>>> a1f77ceab5fe9317d46f796ad157c7ae4e4b40e1

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

<<<<<<< HEAD
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
=======
var MongoClient *mongo.Client

func Connectbase() {
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	clientOption := options.Client().ApplyURI(settings.Mysetting.Get_MongoURL())
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		panic(fmt.Sprintf("MongoDB connect error:%v",err))
	}
	if err =client.Ping(ctx,nil);err!=nil{
		panic(fmt.Sprintf("MongoDB ping failed - is MongoDB running ? error :%v",err))
	}
	fmt.Println("MongoDB connected successfully!")
	MongoClient = client
}
>>>>>>> a1f77ceab5fe9317d46f796ad157c7ae4e4b40e1
