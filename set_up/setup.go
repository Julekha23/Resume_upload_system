package setup

import (
	settings "First_Project/Setting"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
