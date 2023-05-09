package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image-resize/model"
	"log"
	"net/http"
)

func SendRequest(data model.ImageData) {
	url := "http://localhost:8080/image-data"

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Data that was sent...")
	fmt.Println(data)
}
