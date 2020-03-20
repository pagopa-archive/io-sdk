package main

import (
	"bytes"
	"html/template"
)

// Main function for the action
func Main(data map[string]interface{}) map[string]interface{} {

	t := template.Must(template.New("answer").Parse(answerTpl))
	buf := new(bytes.Buffer)
	t.Execute(buf, data)

	entry := Entrypoint{
		URL: data["io-apihost"].(string),
		Key: data["io-apikey"].(string),
	}
	message := Message{
		Dest:     data["dest"].(string),
		Subject:  data["subject"].(string),
		Markdown: data["markdown"].(string),
	}
	m, err := SendMessage(&entry, &message)
	if err != nil {
		m = map[string]interface{}{
			"error": err.Error(),
		}
	}
	return m
}
