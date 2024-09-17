package config

type Config struct {
	ServerAddress string
	// Add other configuration fields as needed
}

func Load() (*Config, error) {
	// Implement configuration loading logic
	return &Config{
		ServerAddress: ":8080",
	}, nil
}
