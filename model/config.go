package model

type Config struct {
	Port string

	// Inherit database config details
	DatabaseConfig
}

type DatabaseConfig struct {
	DBAddr     string
	DBUser     string
	DBPassword string
	DBName     string
}
