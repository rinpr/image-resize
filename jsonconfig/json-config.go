package jsonconfig

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Documentation for parsing json
// https://tutorialedge.net/golang/parsing-json-with-golang/

type ImageConfig struct {
	ImageConfig []Config `json:"image_config"`
}

type Config struct {
	Directory string `json:"directory"`
	Format    string `json:"format"`
	Width     int    `json:"width"`
}

func readJson() (conf ImageConfig) {
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(jsonFile)

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our Users array
	var settings ImageConfig

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'settings' which we defined above
	err = json.Unmarshal(byteValue, &settings)
	if err != nil {
		log.Fatal(err)
	}

	return settings
}

func GetConfig() Config {
	configs := readJson()
	return configs.ImageConfig[0]
}
