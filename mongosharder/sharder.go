package mongosharder

type Sharder interface {
	Shard(Configuration) error
	Reshard(Configuration) error
}
