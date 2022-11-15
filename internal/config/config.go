package config

import "os"

// Config stores configuration
type Config struct {
	// Connection strings
	ServerAddr  string
	RabbitURL   string
	// Rabbit exchange name
	Exchange string
	// Queue names
	QueueBack  string
	// Routing key names
	KeyFront string
	KeyBack  string
}

// New returns configuration variables from the environment.
// These are passed by Docker from the .env file.
func New() *Config {
	return &Config{
		ServerAddr:  getEnv("SERVER_ADDR", "localhost:8080"),
		RabbitURL:   getEnv("RABBIT_URL", "amqp://guest:guest@localhost:5672"),
		Exchange:    getEnv("EXCHANGE", "main_exchange"),
		QueueBack:   getEnv("QUEUE_BACK", "backend_queue"),
		KeyFront:    getEnv("KEY_FRONT", "frontend_key"),
		KeyBack:     getEnv("KEY_BACK", "backend_key"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
