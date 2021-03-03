package models

type Config struct {
	Port string
	RouterConfig
	// Inherit database config details
	DatabaseConfig
}

type DatabaseConfig struct {
	DBAddr     string `required:"true"`
	DBUser     string `required:"true"`
	DBPassword string `required:"true"`
	DBName     string `required:"true"`
}

type JWTConfig struct {
	SecretKey       string `required:"true"`
	Issuer          string `required:"true"`
	ExpirationHours int64  `required:"true"`
}

type RouterConfig struct {
	JWTConfig
}
