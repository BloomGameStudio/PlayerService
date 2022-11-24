package config

import "github.com/spf13/viper"

func Init() {
	ViperInit()

}

func ViperInit() {
	// Initialize Viper with project configuration.

	// Setting Default Values
	viper.SetDefault("DEBUG", true)
	viper.SetDefault("PORT", "1323") // Has to be a string as Echo expects a string
}
