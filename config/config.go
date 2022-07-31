package config

type Config struct {
	Database    DatabaseConfigs
	ServiceHost string
}

type DatabaseConfigs struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func NewConfig() *Config {
	// Hard coding sensitive info for now. Use env variables to replace this step
	return &Config{
		Database: DatabaseConfigs{
			Username: "admin",
			Password: "Password123",
			Host:     "127.0.0.1",
			Port:     3306,
			Database: "company",
		},
		ServiceHost: "localhost:8080",
	}
}
