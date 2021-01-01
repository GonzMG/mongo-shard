package mongoshard

type Sharder interface {
	Shard(Configuration) error
	Reshard(Configuration) error
}
