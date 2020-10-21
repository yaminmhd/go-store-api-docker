package config

import (
	"github.com/spf13/viper"
	"log"
)

func getInt(key string, required bool) int {
	if required {
		checkKey(key)
	}

	return viper.GetInt(key)
}

func getString(key string, required bool) string {
	if required {
		checkKey(key)
	}

	return viper.GetString(key)
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		log.Panicf("Missing key: %s", key)
	}
}