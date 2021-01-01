package mongosharder

type Configuration struct {
	ID             int
	Host           string
	Port           string
	DatabaseName   string
	CollectionName string
	ShardedKey     string
	ShardedNumber  int
}

func NewConfiguration(
	host, port, databaseName, collectionName, shardedKey string,
	shardedNumber int,
) *Configuration {
	newConf := new(Configuration)
	newConf.Host = host
	newConf.Port = port
	newConf.DatabaseName = databaseName
	newConf.CollectionName = collectionName
	newConf.ShardedKey = shardedKey
	newConf.ShardedNumber = shardedNumber
	return newConf
}

func GetConfiguration() Configuration {
	return Configuration{}
}
