package config

import (
	"os"
)

type Config struct {
	RedisConnectionString  string
	RabbitConnectionString string
	RabbitQueueName        string
	NewsApiKey             string
}

func New() *Config {
	return &Config{
		RedisConnectionString:  getEnv("REDIS_URL", ""),
		RabbitConnectionString: getEnv("CLOUDAMQP_URL", ""),
		RabbitQueueName:        getEnv("RABBIT_QUEUE_NAME", ""),
		NewsApiKey:             getEnv("NEWS_API_KEY", ""),
	}
}
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
