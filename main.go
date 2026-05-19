package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//type ImageURL struct {
//	ImageURL string `xml:",chardata"`
//}

func getImage(url string) {
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

	time.Sleep(3 * time.Second)
}

func getFilename() (string) {
	for {
		fmt.Print("Please enter the filename of the Wordpress Image Export XML: ")
		inputReader := bufio.NewReader(os.Stdin)
		inputString, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("file path error: %v\n", err)
		} else {
			return inputString
		}
	}
}

func main() {

	//xmlFile, err := os.ReadFile(getFilename())
	xmlFile, err := os.ReadFile("whitherward.WordPress.2026-05-17.xml")
	if err != nil {
		fmt.Println(fmt.Errorf("xml file error: %v", err))
	}


	xmlBuffer := bytes.NewBuffer(xmlFile)
	decoder := xml.NewDecoder(xmlBuffer)

	foundGUID := false

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "guid" {
				foundGUID = true
				fmt.Printf("found a guid!\n")
			}
		case xml.CharData:
			if foundGUID == true {
				address := strings.TrimSpace(string(t))
				fmt.Printf("Image address: %v\n", address)
				getImage(address)
				foundGUID = false
			}
		}
	}
}