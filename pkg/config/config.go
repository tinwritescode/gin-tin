package config

import "os"

type Config struct {
	ServerAddress string
	JWTSecretKey  string
	// Add other configuration fields as needed
}

func Load() (*Config, error) {
	return &Config{
		ServerAddress: ":8080",
		JWTSecretKey:  os.Getenv("JWT_SECRET_KEY"),
	}, nil
}
