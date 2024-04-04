package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"todo/configs"
)

func NewClient(ctx context.Context, config configs.Config) (db *mongo.Database, err error) {
	var mongoDBURL string
	var isAuth bool
	if config.MongodbUser == "" && config.MongodbPassword == "" {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", config.MongodbHost, config.MongodbPort)
	} else {
		isAuth = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", config.MongodbUser, config.MongodbPassword, config.MongodbHost, config.MongodbPort)
	}

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	authDB := config.MongodbAuth
	if isAuth {
		if authDB == "" {
			authDB = config.MongodbName
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   config.MongodbUser,
			Password:   config.MongodbPassword,
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err)
	}

	return client.Database(config.MongodbName), nil
}
