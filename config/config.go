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
	CountryCodes           [5]string
}

func New() *Config {
	return &Config{
		RedisConnectionString:  getEnv("REDIS_URL", ""),
		RabbitConnectionString: getEnv("CLOUDAMQP_URL", ""),
		RabbitQueueName:        getEnv("RABBIT_QUEUE_NAME", ""),
		NewsApiKey:             getEnv("NEWS_API_KEY", ""),
		Port:                   getEnv("PORT", "5000"),
		MongoDBUri:             getEnv("MONGODB_URI", ""),
		CountryCodes:           [5]string{"hu", "gb", "us", "ca", "de"},
	}
}
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
