package model

type Config struct {
	Port string

	// Inherit database config details
	DatabaseConfig
}

type DatabaseConfig struct {
	DBAddr     string `required:"true"`
	DBUser     string `required:"true"`
	DBPassword string `required:"true"`
	DBName     string `required:"true"`
}
