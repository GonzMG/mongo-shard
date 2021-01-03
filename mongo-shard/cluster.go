package mongoshard

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Cluster struct {
	masterConnection *mongo.Client
	ctx              context.Context
	nodes            []*node
}

type node struct {
	id         int
	connection *mongo.Client
}

// InitCluster creates all the connections with the cluster and stores
// the configuration in the master node
func InitCluster(config *Configuration) *Cluster {
	newCluster := new(Cluster)

	// initialize master connection
	masterConn, err := newMongoClient(config.MasterHost, config.MasterPort)
	if err != nil {
		log.Fatal(err)
	}
	newCluster.masterConnection = masterConn

	// initialize nodes connections
	newCluster.nodes = make([]*node, len(config.NodeHosts))
	for i := range config.NodeHosts {
		nodeConn, err := newMongoClient(config.NodeHosts[i], config.NodePorts[i])
		if err != nil {
			log.Fatal(err)
		}
		newNode := new(node)
		newNode.id = i
		newNode.connection = nodeConn

		newCluster.nodes[i] = newNode
	}

	// save the configuration
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = upsertConfiguration(ctx, config, newCluster.masterConnection)
	if err != nil {
		log.Fatal(err)
	}

	return newCluster
}
