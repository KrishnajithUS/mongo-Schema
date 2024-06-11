package queries

import (
	"context"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BenchmarkMongoWrite(b *testing.B) {
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
	MongoDatabase := client.Database("newDB2")
	User := MongoDatabase.Collection("user")

	b.ResetTimer()
	//startTime := time.Now()
	for i := 0; i < b.N; i++ {
		//startTime = time.Now()
		MongoWrite(User, ctx)
		//readDuration := time.Since(startTime)
		//fmt.Printf("Write operation took %v\n", readDuration)
	}
}

func BenchmarkMongoRead(b *testing.B) {
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
	MongoDatabase := client.Database("newDB2")
	User := MongoDatabase.Collection("user")
	User.Drop(ctx)

	b.ResetTimer()
	//startTime := time.Now()
	for i := 0; i < b.N; i++ {
		//startTime = time.Now()
		MongoRead(User, ctx)
		//readDuration := time.Since(startTime)
		//fmt.Printf("Read operation took %v\n", readDuration)

	}
}

func BenchmarkMongoReadWriteParallel(b *testing.B) {
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
	MongoDatabase := client.Database("newDB2")
	User := MongoDatabase.Collection("user")

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		//startTime := time.Now()
		for pb.Next() {
			//startTime = time.Now()
			MongoRead(User, ctx)
			MongoWrite(User, ctx)
			//readDuration := time.Since(startTime)
			//	fmt.Printf("Read operation took %v\n", readDuration)

		}
	})
}
