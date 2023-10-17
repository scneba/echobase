package config

import "os"

type Config struct {
	ConnectionString string
}

var _configs Config

func SetConfig(conf Config) {
	_configs = conf
}

func GetConfig() Config {
	return _configs
}

func BuildConfigs() Config {

	conf := Config{
		ConnectionString: getEnv("CONNECTION_STRING", "host=localhost port=5432 user=postgres dbname=users password=postgres sslmode=disable"),
	}

	SetConfig(conf)
	return conf
}

// get an env with default values
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
