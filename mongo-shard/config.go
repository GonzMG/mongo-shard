package mongoshard

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Configuration struct {
	ID             string   `bson:"id"`
	MasterHost     string   `bson:"master_host"`
	MasterPort     string   `bson:"master_port"`
	NodeHosts      []string `bson:"node_hosts"`
	NodePorts      []string `bson:"node_ports"`
	DatabaseName   string   `bson:"database"`
	CollectionName string   `bson:"collection"`
	ShardedKey     string   `bson:"shared_key"`
	ShardedNumber  int      `bson:"shared_number"`
}

// NewConfiguration creates a Configuration struct for initializing the sharding service
func NewConfiguration(
	host, port, databaseName, collectionName, shardedKey string,
	shardedNumber int,
	nodeIPs []string,
) *Configuration {
	newConf := new(Configuration)
	newConf.ID = "cluster-config"
	newConf.MasterHost = host
	newConf.MasterPort = port
	newConf.DatabaseName = databaseName
	newConf.CollectionName = collectionName
	newConf.ShardedKey = shardedKey
	newConf.ShardedNumber = shardedNumber

	newConf.NodeHosts, newConf.NodePorts = make([]string, len(nodeIPs)), make([]string, len(nodeIPs))

	for i, nodeIP := range nodeIPs {
		ip := strings.Split(nodeIP, ":")
		newConf.NodeHosts[i] = ip[0]
		newConf.NodePorts[i] = ip[1]
	}

	return newConf
}

func upsertConfiguration(ctx context.Context, config *Configuration, client *mongo.Client) error {
	collection := client.Database(config.DatabaseName).Collection("test")

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"id": config.ID}
	update := bson.M{"$set": &config}

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func GetConfiguration() Configuration {
	return Configuration{}
}
