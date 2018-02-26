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
	currentTime := time.Now().Format("02/01/2006 15:04:05") //This is how you do dates in go.. they are magic numbers
	fileLocation = strings.Replace(fileLocation, "\\", "/", -1)

	log.Println("Version 0.3")

	payloadRaw := "{\"@type\":\"MessageCard\",\"@context\":\"http://schema.org/extensions\",\"Summary\":\"Testing\",\"themeColor\":\"00ff00\",\"sections\":[{\"activityTitle\":\"New Data has arrived on the FTP Server\",\"facts\":[{\"name\":\"User:\",\"value\":\"XXXX\"},{\"name\":\"File:\",\"value\":\"YYYY\"},{\"name\":\"Size:\",\"value\":\"ZZZZ\"},{\"name\":\"Date:\",\"value\":\"TTTT\"}]}]}"
	payloadRaw = strings.Replace(payloadRaw, "XXXX", name, 1)
	payloadRaw = strings.Replace(payloadRaw, "YYYY", fileLocation, 1)
	payloadRaw = strings.Replace(payloadRaw, "ZZZZ", fileSize, 1)
	payloadRaw = strings.Replace(payloadRaw, "TTTT", currentTime, 1)

	log.Println(payloadRaw)

	url := "https://outlook.office.com/webhook/686a7071-f9ea-4fca-865e-45d3c8f674b2@74bc964f-1d6e-4ec4-8ec5-4c77c63906a9/IncomingWebhook/3c2eadeb7b734d7cb5ef477cc1f864bf/b4f61c5c-830a-4995-b490-8ed756ea5fc0"

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

