package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"url-shorter/internal/generatorShortUrl"
	"url-shorter/pkg/errorsHandling"
)

type Urls struct {
	Url   string `bson:"url"`
	Short string `bson:"short"`
}

const (
	location                  = "internal.database.mongodb"
	warningDisconnectDatabase = "warning: disconnect database"

	shortUrlKey = "short"
	urlKey      = "url"
)

func (s *MongoStorage) Disconnect() {
	err := s.Client.Disconnect(Ctx)
	if err != nil {
		log.Println(warningDisconnectDatabase)
	}
}

func (s *MongoStorage) GetShortUrl(url string) (shortUrl string, err error) {
	if urlExists(s.Collection, url) {
		shortUrl, err = getShortUrl(s.Collection, url)
	} else {
		shortUrl, err = saveNewUrl(s.Collection, url)
	}
	return
}

func urlExists(collection *mongo.Collection, url string) bool {
	return exists(collection, urlKey, url)
}

func exists(collection *mongo.Collection, key, value string) bool {
	urls, err := getUrls(collection, key, value)
	return err == nil && urls.Url != "" && urls.Short != ""
}

func saveNewUrl(collection *mongo.Collection, url string) (shortUrl string, err error) {
	shortUrl = generatorShortUrl.GenerateShortUrl()
	_, err = collection.InsertOne(Ctx, Urls{url, shortUrl})
	if err != nil {
		err = errorsHandling.ErrFormat(location, "saveNewUrl", err)
	}
	return
}

func getShortUrl(collection *mongo.Collection, url string) (shortUrl string, err error) {
	shortUrl, err = findShortUrl(collection, url)
	return
}

func findShortUrl(collection *mongo.Collection, url string) (string, error) {
	urls, err := getUrls(collection, urlKey, url)
	return urls.Short, err
}

func getUrls(collection *mongo.Collection, key, value string) (urls Urls, err error) {
	filter := makeFilter(key, value)
	raw, err := findRawUrls(collection, filter)
	if err != nil {
		err = errorsHandling.ErrFormat(location, "getUrls", err)
		return
	}
	urls, err = decodeRaw(raw)
	return
}

func makeFilter(key, value string) bson.M {
	return bson.M{
		key: value,
	}
}

func findRawUrls(collection *mongo.Collection, filter bson.M) (raw bson.Raw, err error) {
	raw, err = collection.FindOne(Ctx, filter).DecodeBytes()
	if err != nil {
		err = errorsHandling.ErrFormat(location, "findRawUrls", err)
	}
	return
}

func decodeRaw(raw bson.Raw) (urls Urls, err error) {
	values, err := raw.Values()
	if err != nil {
		err = errorsHandling.ErrFormat(location, "decodeRaw", err)
	}
	urls = valuesFormat(values)
	return
}

func valuesFormat(values []bson.RawValue) Urls {
	return Urls{
		Url:   values[1].StringValue(),
		Short: values[2].StringValue(),
	}
}

func (s *MongoStorage) GetUrl(shortUrl string) (string, error) {
	urls, err := getUrls(s.Collection, shortUrlKey, shortUrl)
	return urls.Url, err
}
