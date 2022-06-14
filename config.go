package main

import (
	"log"

	"github.com/spf13/viper"
)

// Configurations : Port number and MySQL connection string
type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

// App config variable accessed by other files
var AppConfig *Config

// Load configurations with Viper from the config.json file and assign values into AppConfig variable
func LoadAppConfig() {
	log.Println("Loading server configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json") // Here a JSON file but could be a yaml file or env.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err) // Fatal ->Print + Exit(1)
	}
	err = viper.Unmarshal(&AppConfig) // JSON to Struct
	if err != nil {
		log.Fatal(err)
	}
}
