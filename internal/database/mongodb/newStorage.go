package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"url-shorter/internal/config"
)

type Storage interface {
	Disconnect()
	GetShortUrl(string) (string, error)
	GetUrl(string) (string, error)
}

type MongoStorage struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

const (
	ErrorConnectionClient = "failed to connect to mongodb"
)

var (
	Ctx = context.Background()
)

func MustNewStorage(cfg config.MongodbConfig) Storage {
	client := connectToMongodb(Ctx, cfg)
	collection := getCollection(client, cfg.Database, cfg.Collection)
	var storage = MongoStorage{
		Client:     client,
		Collection: collection,
	}
	return &storage
}

func connectToMongodb(ctx context.Context, cfg config.MongodbConfig) *mongo.Client {
	optionsClient := getOptionsClient(cfg.Uri)
	client, err := mongo.Connect(ctx, optionsClient)
	if err != nil {
		log.Fatal(ErrorConnectionClient)
	}
	return client
}

func getOptionsClient(uri string) *options.ClientOptions {
	return options.Client().ApplyURI(uri)
}

func getCollection(client *mongo.Client, database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
