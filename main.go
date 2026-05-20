package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	LO_RATE int = 1
	HI_RATE int = 60
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

	if err := os.WriteFile("./downloads/" + fileName, data, 0666); err != nil {
		fmt.Printf("file write error: %v\n", err)
	}

	fmt.Println("File downloaded successfully")

}

func getFilename(reader *bufio.Reader) (string) {

	fmt.Print("Please enter the filename of the Wordpress Image Export XML, or the full path if file is not in the current working directory: ")
	inputString, err := reader.ReadString('\n')
	fmt.Printf("%v", inputString)

	if err != nil {
		fmt.Printf("file path error: %v\n", err)
	}
	return strings.Trim(inputString, "\n")
}

func main() {

	//get file and rate info
	downRate := 0
	var xmlPath string
	var xmlFile []byte
	inputReader := bufio.NewReader(os.Stdin)
	
	for {
		xmlPath = getFilename(inputReader)
		t, err := os.ReadFile(xmlPath)
		if err != nil {
			fmt.Println(fmt.Errorf("xml file error: %v", err))
			continue
		}
		xmlFile = t
		break
	}

	for {
		fmt.Print("Enter a time between downloads in seconds (min 1, max 60): ")
		_, err := fmt.Scan(&downRate)
		if err != nil {
			fmt.Printf("input error: %v\n", err)
			continue
		}
		
		if downRate < LO_RATE {
			fmt.Println("invalid rate")
			continue
		} else if downRate > HI_RATE {
			fmt.Println("invalid rate")
			continue
		} else {
			fmt.Printf("seconds of delay: %v\n", downRate)
			break
		}
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
				time.Sleep(time.Duration(downRate) * time.Second)
				getImage(address)
				foundGUID = false
			}
		}
	}
}