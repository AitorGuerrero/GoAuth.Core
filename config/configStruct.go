package config

type SqlDbConfig struct {
	Name string
	UserName string
	Password string
	Host string
}

type KiteServiceConfig struct {
	Name string
	Version string
	Port int
}

type Config struct {
	SqlDbConfig
	KiteServiceConfig
}
