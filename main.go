package main

import mongoshard "mongo-shard/mongo-shard"

func main() {
	conf := mongoshard.NewConfiguration("127.0.0.1", "27017", "test", "main", 2, []string{"127.0.0.1:27017"})
	mongoshard.InitCluster(conf)
}
