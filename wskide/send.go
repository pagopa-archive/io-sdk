package wskide

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	sendDestArg     = sendCmd.Arg("recipient", "Message Recipient (fiscal code)").String()
	sendSubjectArg  = sendCmd.Arg("subject", "Message Subject").String()
	sendMarkdownArg = sendCmd.Arg("markdown", "Message Markdown").String()
	sendKeyArg      = sendCmd.Arg("key", "API Key").String()
)

// SendMessage send message to IO
func SendMessage(subject, markdown, dest, key string) (err error) {

	url := APIHOST + "/messages"
	t := time.Now().UTC().Format("2006-01-02T15:04:05.070Z")
	message := `{
		"time_to_live": 3600,
		"content": {
			"subject": "%s",
			"markdown": "%s",
			"due_date": "%s"
		},
		"fiscal_code": "%s"
	}`
	message = fmt.Sprintf(message,
		subject,
		markdown,
		t,
		dest,
	)
	//fmt.Println("URL:>", url)
	var jsonStr = []byte(message)
	size := len(jsonStr)
	//fmt.Println(message)
	//fmt.Println("Content-Lenght", size)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Ocp-Apim-Subscription-Key", key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Lenght", string(size))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return err

}
