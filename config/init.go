package config

import "github.com/spf13/viper"

func Init() {
	ViperInit()

}

func ViperInit() {
	// Initialize Viper with project configuration.

	// Setting Default Values
	viper.SetDefault("DEBUG", true)
}
