package config

import "fmt"

type DatabaseConfig struct {
	host     string
	username string
	password string
	name     string
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		host:     getString("DB_HOST", true),
		name:     getString("DB_NAME", true),
		username: getString("DB_USER", true),
		password: getString("DB_PASSWORD", true),
	}
}

func (dc DatabaseConfig) ConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dc.username, dc.password, dc.host, dc.name)
}

func (dc DatabaseConfig) DatabaseName() string {
	return dc.name
}
