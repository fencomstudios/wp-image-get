package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := "https://whitherward.com/wp-content/uploads/2017/02/ww_photo_2_dark.jpg"
	urlSegments := strings.Split(url, "/")
	fileName := urlSegments[len(urlSegments) - 1]

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("request error: %v", err))
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("response error: %v", err))
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("data read error: %v", err))
	}

	fmt.Println("Content type is: " + http.DetectContentType(data))

	fileErr := os.WriteFile("./downloads/" + fileName, data, 0666)
	if fileErr != nil {
		log.Fatal(err)
	}
}