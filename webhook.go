package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
        "time"
)

func main() {

	name := os.Getenv("VIRTUSER")
	fileLocation := os.Getenv("SSHUPLOADFILE")
	fileSize := os.Getenv("SSHUPLOADBYTES")
	currentTime := time.Now().Format(time.RFC3339)
	fileLocation = strings.Replace(fileLocation, "\\", "/", -1)

	log.Println("Version 1.0")

	payloadRaw := "{\"@type\":\"MessageCard\",\"@context\":\"http://schema.org/extensions\",\"Summary\":\"Testing\",\"themeColor\":\"00ff00\",\"sections\":[{\"activityTitle\":\"New Data has arrived on the FTP Server\",\"facts\":[{\"name\":\"User:\",\"value\":\"XXXX\"},{\"name\":\"File:\",\"value\":\"YYYY\"},{\"name\":\"Size:\",\"value\":\"ZZZZ\"},{\"name\":\"Date:\",\"value\":\"TTTT\"}]}]}"

	payloadRaw = strings.Replace(payloadRaw, "XXXX", name, 1)
	payloadRaw = strings.Replace(payloadRaw, "YYYY", fileLocation, 1)
	payloadRaw = strings.Replace(payloadRaw, "ZZZZ", fileSize, 1)

	log.Println(payloadRaw)

	url := "https://outlook.office.com/webhook/......"

	log.Println("Posting..")

	payload := strings.NewReader(payloadRaw)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	ioutil.ReadAll(res.Body)

	log.Println(res)
}

