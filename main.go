package main

import (
	"context"
	"fmt"
	"log"

	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email,omitempty"`
	Phone string             `bson:"phone"`
	Age   int                `bson:"age"`
	Sex   string             `bson:"sex,omitempty"`
}

// var ctx * context.Context
func MongoWrite(users *mongo.Collection, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	log.Println("Mongo write Started at: ", start)
	// res, err := users.InsertMany(context.TODO(), []interface{}{
	// 	bson.D{{Key: "name", Value: "A"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
	// 	bson.D{{Key: "name", Value: "B"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
	// 	bson.D{{Key: "name", Value: "C"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	// 	bson.D{{Key: "name", Value: "D"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
	// 	bson.D{{Key: "name", Value: "E"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
	// 	bson.D{{Key: "name", Value: "F"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	// 	bson.D{{Key: "name", Value: "G"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
	// 	bson.D{{Key: "name", Value: "h"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
	// 	bson.D{{Key: "name", Value: "i"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	// 	bson.D{{Key: "name", Value: "test4"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	// })
	res, err := users.InsertOne(context.TODO(), []interface{}{
		bson.D{{Key: "name", Value: "test5"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	})
	//res, err := users.InsertOne(context.TODO(), bson.D{{Key: "name", Value: "test"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}})
	log.Println("Insert Result :", res.InsertedID)
	log.Println("Mongo Write finished at : ", start)
	log.Println("Mongo Write time taken : ", time.Since(start))
	if err != nil {
		log.Println("Error", err.Error())
	}

}

func MongoRead(users *mongo.Collection, wg *sync.WaitGroup) {
	defer wg.Done()
	var userData User
	start := time.Now()
	log.Println("Mongo Reading Started at: ", start)
	userCursor := users.FindOne(context.TODO(), bson.D{{Key: "name", Value: "test4"}})
	log.Println("Mongo Reading finished at : ", start)
	log.Println("Mongo Reading time taken : ", time.Since(start))
	err := userCursor.Decode(&userData)
	fmt.Println(userData, err)
}

func main() {

	if len(os.Args) < 1 {
		panic("Usage: program [readPreference]")
	}

	var uri string
	if len(os.Args) == 2 {
		readPreference := os.Args[1]
		uri = fmt.Sprintf("mongodb://sa:Password123@127.0.10.1:27017,127.0.10.2:27017,127.0.10.3:27017/?replicaSet=rs0&readPreference=%s", readPreference)
	} else {
		uri = "mongodb://sa:Password123@127.0.10.1:27017,127.0.10.2:27017,127.0.10.3:27017/?replicaSet=rs0"
	}
	fmt.Println(uri)
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	DB := client.Database("newDB2")
	users := DB.Collection("user")

	var wg sync.WaitGroup

	wg.Add(2)
	go MongoRead(users, &wg)
	go MongoWrite(users, &wg)

	wg.Wait()
	//	elapsedRead := time.Since(startRead)
	//	log.Printf("MongoRead took %s", elapsedRead)
	// primary, seondary, err := getReplicaSetStatus(client, ctx)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("Primary Replicas : ")

	// fmt.Println(primary)

	// log.Printf("Secondary Replicas : ")
	// for i := range seondary {
	// 	fmt.Println(seondary[i])
	// }
	//log.Printf("Read Servers: %v", readServers)
	//}
	// file, err := os.OpenFile("events.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatalf("Failed to open log file: %s", err)
	// }
	// defer file.Close()
	// // Create a logger that writes to the file
	// logger := log.New(file, "", log.LstdFlags)
	// // Log the servers used for write and read operations
	// logger.Println("Started operations were performed on the following servers:")
	// for server := range cmdMonitorStart {
	// 	logger.Println(cmdMonitorStart[server])
	// }
	// logger.Println("Success operations were performed on the following servers:")
	// for server := range cmdMonitorSucceed {
	// 	logger.Println(cmdMonitorSucceed[server])
	// }
	// logger.Println("Failed operations were performed on the following servers:")
	// for server := range cmdMonitorFailed {
	// 	logger.Println(cmdMonitorSucceed[server])
	// }
	// for i := range eventArray {
	// 	logger.Println("Writing Events", eventArray[i])
	// }
}
