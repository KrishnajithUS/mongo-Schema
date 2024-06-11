package main

import (
	"context"
	"fmt"
	"log"
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
	Phone string	`bson:"phone"`
	Age   int                `bson:"age"`
	Sex string `bson:"sex,omitempty"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
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
	
	var userData []User
	// userCursor.All(ctx, &userData)
	// //log.Println(userData)
	// for _, user := range userData {
	// 	fmt.Println(user)
	// }
	log.Println("Entering flexible data before schema altering")
	//adding flexible data
	result, err := users.InsertMany(ctx, []interface{}{
		bson.D{{Key: "name", Value: "gopu"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
		bson.D{{Key: "name", Value: "hari"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
		bson.D{{Key: "name", Value: "indu"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	})
	log.Println("Inserted flexible data before schema altering")

	//log.Println("result", result, err)
	//printing all data
	// userCursor.All(ctx, &userData)
	// //log.Println(userData)
	// for _, user := range userData {
	// 	fmt.Println(user)
	// }
	log.Println("Adding validator for schema altering")

	//adding validator
	validator := bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"name", "phone", "age"},
			"properties": bson.M{
				"name": bson.M{
					"bsonType":    "string",
					"description": "must be a string and is required",
				},
				"phone": bson.M{
					"bsonType":    "string",
					"description": "must be an string and is required",
				},
				"age": bson.M{
					"bsonType":    "int",
					"description": "must be a int and is required",
				},
			},
		},
	}

	collModCmd := bson.D{
		{Key: "collMod", Value: "user"},
		{Key: "validator", Value: validator},
	}
	log.Println("completed adding validator for schema altering")
	err = DB.RunCommand(ctx, collModCmd).Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Collection user modified with validation\n")
	//adding flexible data
	log.Println("Entering flexible data after schema altering")

	result2, err := users.InsertMany(ctx, []interface{}{
		bson.D{{Key: "name", Value: "janu"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
		bson.D{{Key: "name", Value: "kira"}, {Key: "phone", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
		bson.D{{Key: "username", Value: "linto"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "M"}},
	})
	if err!=nil{
		log.Println("Error:",err.Error())
	}

	//adding required data
	log.Println("Entering required format data after schema altering")

	result3, err := users.InsertMany(ctx, []interface{}{
		bson.D{{Key: "name", Value: "manu"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}, {Key: "phone", Value: "934252345234"}},
		bson.D{{Key: "name", Value: "niru"}, {Key: "phone", Value: "234787654"}, {Key: "age", Value: 47}},
		bson.D{{Key: "name", Value: "omana"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "M"}, {Key: "age", Value: 23}},
	})
	log.Println("completed: Printing results")

	log.Println(result, result2, result3)
	//printing all data
	log.Println("All collections")
	userCursor, _ := users.Find(ctx, bson.D{})
	userCursor.All(ctx, &userData)
	//log.Println(userData)
	for _, user := range userData {
		fmt.Println(user)
	}

}

// //start-session
// wc := writeconcern.Majority()
// txnOptions := options.Transaction().SetWriteConcern(wc)

// // Starts a session on the client
// session, err := client.StartSession()
// if err != nil {
// 	panic(err)
// }
// // Defers ending the session after the transaction is committed or ended
// defer session.EndSession(context.TODO())
// // 	 Inserts multiple documents into a collection within a transaction,
// // 	 then commits or ends the transaction
// result, err := session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
// 	result, err := users.InsertMany(ctx, []interface{}{
// 		bson.D{{"name", "The Bluest Eye"}, {"email", "Toni Morrison"},{"age",24}},
// 		bson.D{{"name", "Sula"}, {"email", "Toni Morrison"},{"age",47}},
// 		bson.D{{"name", "Song of Solomon"}, {"email", "Toni Morrison"},{"age",14}},
// 	})
// 	return result, err
// }, txnOptions)
// // end-session

// fmt.Printf("Inserted _id values: %v\n", result)
