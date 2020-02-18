package store


// Config ...
type Config string {
	DatabaseURL string `toml:"database_url"`
}


// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}

