package main

import (
	"context"
	"log"
	"time"
	"main/queries"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")) // connectionstring
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	DB := client.Database("newDB")
	users := DB.Collection("user")
	queries.MongoWrite(users, ctx)
	queries.MongoRead(users, ctx)

	//============ Code for write data
	//log.Println("Entering flexible data before schema altering")
	//adding flexible data
	// result, err := users.InsertMany(ctx, []interface{}{
	// 	bson.D{{Key: "name", Value: "A"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
	// 	bson.D{{Key: "name", Value: "B"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
	// 	bson.D{{Key: "name", Value: "C"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	// 	bson.D{{Key: "name", Value: "D"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
	// 	bson.D{{Key: "name", Value: "E"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
	// 	bson.D{{Key: "name", Value: "F"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	// 	bson.D{{Key: "name", Value: "G"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
	// 	bson.D{{Key: "name", Value: "h"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
	// 	bson.D{{Key: "name", Value: "i"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	// 	bson.D{{Key: "name", Value: "j"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	// })
	// log.Println("Inserted flexible data before schema altering")

	// log.Println(result)

	// //=========== Code for read data
	// log.Println("All collections")
	// userCursor, _ := users.Find(ctx, bson.D{})
	// userCursor.All(ctx, &userData)
	// //log.Println(userData)
	// for _, user := range userData {
	// 	fmt.Println(user)
	// }

}
