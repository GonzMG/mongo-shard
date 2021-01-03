package main

import mongoshard "mongo-shard/mongo-shard"

func main() {
	conf := mongoshard.NewConfiguration("172.17.0.2", "27017", "test", "main", "id", 2, []string{"172.17.0.2:27017"})
	mongoshard.InitCluster(conf)
}
