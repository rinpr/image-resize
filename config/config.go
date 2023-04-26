package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func read(config string) (name string, number int) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	name = viper.GetString(config)
	number = viper.GetInt(config)
	return
}

func GetHeight() (height int) {
	_, height = read("height")
	return
}

func GetWidth() (width int) {
	_, width = read("width")
	return
}

func GetName() (name string) {
	name, _ = read("FileName")
	return
}

// https://medium.com/@bnprashanth256/reading-configuration-files-and-environment-variables-in-go-golang-c2607f912b63
