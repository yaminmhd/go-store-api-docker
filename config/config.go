package config

import (
	"github.com/spf13/viper"
)

type config struct {
	env                string
	port               int
	logLevel           string
	dbConfig           DatabaseConfig
}

var appConfig config

func Load(){
	viper.AutomaticEnv()
	viper.SetConfigName("")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	appConfig = config{
		env:                getString("ENVIRONMENT", true),
		port:               getInt("APP_PORT", true),
		logLevel:           getString("LOG_LEVEL", true),
		dbConfig:           getDatabaseConfig(),
	}
}

func Environment() string {
	return appConfig.env
}

func Port() int {
	return appConfig.port
}

func LogLevel() string {
	return appConfig.logLevel
}

func Database() DatabaseConfig {
	return appConfig.dbConfig
}