package mongosharder

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection(config *Configuration) {

}

func newMongoClient(config *Configuration) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
