package main

import "go-sharder/mongoshard"

func main() {
	c := mongoshard.NewConfiguration("127.0.0.1", "27100", "test", "main", "id", 2)
	mongoshard.InitConnection(c)
}
