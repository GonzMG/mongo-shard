package mongoshard

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Cluster struct {
	masterConnection *mongo.Client
	router           map[int]*node
	config           *Configuration
}

type node struct {
	id         int
	connection *mongo.Client
}

// InitCluster creates all the connections with the cluster and stores
// the configuration in the master node
func InitCluster(config *Configuration) *Cluster {
	newCluster := new(Cluster)

	newCluster.config = config

	// initialize master connection
	masterConn, err := newMongoClient(config.MasterHost, config.MasterPort)
	if err != nil {
		log.Fatal(err)
	}
	newCluster.masterConnection = masterConn

	// initialize nodes connections
	nodes := make([]*node, len(config.NodeHosts))
	for i := range config.NodeHosts {
		nodeConn, err := newMongoClient(config.NodeHosts[i], config.NodePorts[i])
		if err != nil {
			log.Fatal(err)
		}
		newNode := new(node)
		newNode.id = i
		newNode.connection = nodeConn

		nodes[i] = newNode
	}

	// create the router map
	router := make(map[int]*node)
	route(nodes, router, config.ShardedNumber)
	newCluster.router = router

	// save the configuration in the specified database into config collection
	err = upsertClusterConfiguration(config, newCluster.masterConnection)
	if err != nil {
		log.Fatal(err)
	}

	return newCluster
}

// route fills the router specifying where will be located each sharded collection
// distributed in the different nodes
//
// the router map will be completed in order, the first sharded collection will be
// in the first node, the N sharded collection in the N node
//
// if the quantity of nodes is smaller than shardedQuantity, there will be nodes
// with more than one sharded collection
func route(nodes []*node, router map[int]*node, shardQuantity int) {
	var idShard, nodeIndex int
	if len(nodes) == 0 {
		return
	}
	for len(router) < shardQuantity {
		if nodeIndex == len(nodes) {
			nodeIndex = 0
		}
		router[idShard] = nodes[nodeIndex]
		idShard++
		nodeIndex++
	}
}
