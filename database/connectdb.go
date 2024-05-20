package database

import (
	"context"
	// "gin-temp-app/constant"
	"log"
	"os"

	// "os"
	"time"

	"gin-temp-app/types"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type manager struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	GetSingleRecordByUsername(string,string) *types.User
	GetNotes(primitive.ObjectID,string) []types.Note
	GetSingleNoteById(string, string) *types.Note
	UpdateNote(map[string]interface{}, string,string) error
	GetUserByEmailPassword(string,string,string) *types.User
	
}

func ConnectDb() {
	uri := os.Getenv("DatabaseURL")
	// if uri == "" {
	// 	uri = constant.MDBUri
	// }
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    
	if err != nil {
		ConnectDb()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// if coll != nil {
	// 	log.Println("Unable to initialize database connectors. Retrying...")
	// 	ConnectDb()
	// }
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("Unable to connect to the database. Retrying...")
		// ConnectDb()
	}
	log.Println("Successfully connected to the database at %s", uri)

	Mgr = &manager{connection: client, ctx: ctx, cancel: cancel}
}

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}