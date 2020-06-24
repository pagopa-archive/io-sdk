package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Entrypoint is the entrypoint to send messages
type Entrypoint struct {
	URL string
	Key string
}

// Message is the message
type Message struct {
	FiscalCode string `json:"fiscal_code"`
	Subject    string `json:"subject"`
	Markdown   string `json:"markdown"`
	//DueDate  string
}

// SendMessage send message to IO
func SendMessage(entry *Entrypoint, msg *Message) (map[string]interface{}, error) {
	//if msg.DueDate == "" {
	//	msg.DueDate = time.Now().UTC().Format("2006-01-02T15:04:05.070Z")
	//}
	t := template.Must(template.New("message").Parse(messageTpl))
	buf := new(bytes.Buffer)
	t.Execute(buf, msg)

	//fmt.Println("Content-Lenght", size)
	req, err := http.NewRequest("POST", entry.URL, buf)
	req.Header.Set("Ocp-Apim-Subscription-Key", entry.Key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", strconv.Itoa(buf.Len()))
	//log.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	json.Unmarshal(body, &res)
	return res, nil
}
