package queries

import (
	"context"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var eventArray []*event.ServerDescriptionChangedEvent
var insertCount int = 0

// var srvMonitor *event.ServerMonitor = &event.ServerMonitor{
// 	ServerDescriptionChanged: func(e *event.ServerDescriptionChangedEvent) {
// 		eventArray = append(eventArray, e)
// 	},
// }

var cmdMonitor *event.CommandMonitor = &event.CommandMonitor{
	Started: func(_ context.Context, evt *event.CommandStartedEvent) {
		// log.Printf("Started command: %s %s\n", evt.CommandName, evt.Command)
	},
	Succeeded: func(_ context.Context, evt *event.CommandSucceededEvent) {
		if evt.CommandName == "insert"{
			insertCount += 1
		}
		// log.Printf("Succeeded command: %s %v\n", evt.CommandName, evt.Reply)
	},
	Failed: func(_ context.Context, evt *event.CommandFailedEvent) {
		// log.Printf("Failed command: %s %v\n", evt.CommandName, evt.Failure)
	},
}

func logEvents() {
	log.Println("Total Insertions :", insertCount)
}

func BenchmarkMongoWrite(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI("mongodb://sa:Password123@127.0.10.1:27017,127.0.10.2:27017,127.0.10.3:27017/?replicaSet=rs0").SetMonitor(cmdMonitor)
	client, err := mongo.Connect(context.TODO(), clientOpts)

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
	logEvents()


}

// func BenchmarkMongoRead(b *testing.B) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	for i := range eventArray {
// 		log.Println("Writing Events", eventArray[i])
// 	}
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://sa:Password123@127.0.10.1:27017,127.0.10.2:27017,127.0.10.3:27017/?replicaSet=rs0"))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer func() {
// 		if err = client.Disconnect(context.TODO()); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()
// 	MongoDatabase := client.Database("newDB2")
// 	User := MongoDatabase.Collection("user")

// 	b.ResetTimer()
// 	//startTime := time.Now()
// 	for i := 0; i < b.N; i++ {
// 		//startTime = time.Now()
// 		MongoRead(User, ctx)
// 		//readDuration := time.Since(startTime)
// 		//fmt.Printf("Read operation took %v\n", readDuration)

// 	}

// }

// func BenchmarkMongoReadWriteParallel(b *testing.B) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://sa:Password123@127.0.10.1:27017,127.0.10.2:27017,127.0.10.3:27017/?replicaSet=rs0").SetServerMonitor(srvMonitor))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer func() {
// 		if err = client.Disconnect(context.TODO()); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()
// 	MongoDatabase := client.Database("newDB2")
// 	User := MongoDatabase.Collection("user")

// 	b.ResetTimer()
// 	b.RunParallel(func(pb *testing.PB) {
// 		//startTime := time.Now()
// 		for pb.Next() {
// 			//startTime = time.Now()
// 			MongoRead(User, ctx)
// 			MongoWrite(User, ctx)
// 			//readDuration := time.Since(startTime)
// 			//	fmt.Printf("Read operation took %v\n", readDuration)

// 		}
// 	})

// }
