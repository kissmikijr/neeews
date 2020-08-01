package config

import (
	"os"
)

type Config struct {
	RedisConnectionString  string
	RabbitConnectionString string
	RabbitQueueName        string
	NewsApiKey             string
	Port                   string
	MongoDBUri             string
	CountryCodes           [3]string
	WorkerToken            string
	HostName               string
}

func New() *Config {
	return &Config{
		RedisConnectionString: getEnv("REDIS_URL", ""),
		NewsApiKey:            getEnv("NEWS_API_KEY", ""),
		Port:                  getEnv("PORT", "5000"),
		MongoDBUri:            getEnv("MONGODB_URI", ""),
		CountryCodes:          [3]string{"hu", "gb", "us"},
		WorkerToken:           getEnv("WORKER_TOKEN", ""),
		HostName:              getEnv("HOST_NAME", "http://localhost:5000"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
