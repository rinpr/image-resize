package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	ImgSize  ImageConfigurations
	FileName string
}

type ImageConfigurations struct {
	Height int
	Width  int
}

func read(config string) (name string, number int) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
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
