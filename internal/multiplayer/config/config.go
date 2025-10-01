package config

// Config for the multiplayer mode
type Config struct {
}

func NewConfig() (*Config, error) {
	return &Config{}, nil
}
