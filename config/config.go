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
	CountryCodes           [3]string
	WorkerToken            string
	HostName               string
}

func New() *Config {
	return &Config{
		RedisConnectionString: getEnv("REDIS_URL", "redis://rediscloud:NxAJR8SP2gjyNdRswDT04BszVT8EtLm3@redis-10177.c3.eu-west-1-2.ec2.cloud.redislabs.com:10177"),
		NewsApiKey:            getEnv("NEWS_API_KEY", "e1750e47dd844d54ac301a7f99c8cdf5"),
		Port:                  getEnv("PORT", "5000"),
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
