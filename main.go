package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type ImageURL struct {
	ImageURL string `xml:",chardata"`
}

func main() {
	url := "https://whitherward.com/wp-content/uploads/2017/02/ww_photo_2_dark.jpg"
	urlSegments := strings.Split(url, "/")
	fileName := urlSegments[len(urlSegments) - 1]

	testData := `
	<rss>
		<channel>
			<item>
				<guid>http://whitherward.com/wp-content/uploads/2017/02/home-1.jpg</guid>
			</item>
		</channel>
	</rss>
	`

	//xmlFile, err := os.ReadFile("whitherward.WordPress.2026-05-17.xml")
	//if err != nil {
	//	fmt.Println(fmt.Errorf("xml file error: %v", err))
	//}

	//var urlList []ImageURL
	testDataReader := strings.NewReader(testData)
	decoder := xml.NewDecoder(testDataReader)
	
	//if err := xml.Unmarshal([]byte(testData), &urlList); err != nil {
	//	fmt.Println(fmt.Errorf("xml file error: %v", err))
	//}

	for {
		token, err := decoder.Token()
		if err != nil {
			fmt.Println(fmt.Errorf("%v", err))
			break
		}
		if err == io.EOF {
			fmt.Println(fmt.Errorf("%v"), err)
			break
		}
		
		if c, ok := token.(xml.StartElement); ok {
			if c.Name.Local == "guid" {
				
			}
		}
	}

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