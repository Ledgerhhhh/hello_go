package myconfig

var GConfig = new(GlobalConfig)

type GlobalConfig struct {
	WsConfig      WsConfig
	RedisConfig   RedisConfig
	MysqlConfig   MysqlConfig
	PulsarConfig  PulsarConfig
	ElasticConfig ElasticConfig
	NsqdConfig    NsqdConfig
	KafkaConfig   KafkaConfig
}

type KafkaConfig struct {
	Host string
	Port int
}

type WsConfig struct {
	IP   string
	Port int
}

type RedisConfig struct {
	IP       string
	Port     int
	Network  string
	Password string
	DB       int
}

type MysqlConfig struct {
	Dsn string
}

type PulsarConfig struct {
	BrokerURL        string
	Topic            string
	SubscriptionName string
}

type ElasticConfig struct {
	Host string
	Port int
}

type NsqdConfig struct {
	Host     string
	Port     int
	Topic    string
	Channel  string
	Channel2 string
}
