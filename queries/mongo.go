package queries

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email,omitempty"`
	Phone string             `bson:"phone"`
	Age   int                `bson:"age"`
	Sex   string             `bson:"sex,omitempty"`
}

func MongoWrite(users *mongo.Collection, ctx context.Context) {
	_, err := users.InsertMany(ctx, []interface{}{
		bson.D{{Key: "name", Value: "A"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
		bson.D{{Key: "name", Value: "B"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
		bson.D{{Key: "name", Value: "C"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
		bson.D{{Key: "name", Value: "D"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
		bson.D{{Key: "name", Value: "E"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
		bson.D{{Key: "name", Value: "F"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
		bson.D{{Key: "name", Value: "G"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 24}},
		bson.D{{Key: "name", Value: "h"}, {Key: "email", Value: "Toni Morrison"}, {Key: "age", Value: 47}},
		bson.D{{Key: "name", Value: "i"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
		bson.D{{Key: "name", Value: "j"}, {Key: "phone", Value: "934252345234"}, {Key: "sex", Value: "F"}},
	})
	if err != nil {
		log.Println("Error",err.Error())
	}
}

func MongoRead(users *mongo.Collection, ctx context.Context) {
	var userData []User

	userCursor, _ := users.Find(ctx, bson.D{{Key:"name",Value: "j"}})
	userCursor.All(ctx, &userData)
}
