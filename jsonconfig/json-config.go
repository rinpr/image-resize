package jsonconfig

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type ImageConfig struct {
	ImageConfig []Config `json:"image_config"`
}

type Config struct {
	Filename string `json:"file_name"`
	Format   string `json:"format"`
	Width    int    `json:"width"`
}

func readJson() (conf ImageConfig) {
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Opened config.json")
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

	// we iterate through every user within our settings array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	//for i := 0; i < len(settings.ImageConfig); i++ {
	//	fmt.Println("File format: " + settings.ImageConfig[i].Format)
	//	fmt.Println("Width: " + strconv.Itoa(settings.ImageConfig[i].Width))
	//}
	return settings
}

func GetConfig(index int) (image Config) {
	configs := readJson()
	image = Config{configs.ImageConfig[index].Filename, configs.ImageConfig[index].Format, configs.ImageConfig[index].Width}
	return
}
