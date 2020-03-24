package main

import (
	"bytes"
	"html/template"
)

func mkMap(k, v string) map[string]interface{} {
	m := make(map[string]interface{})
	m[k] = v
	return m
}

// Main function for the action
func Main(data map[string]interface{}) map[string]interface{} {

	t := template.Must(template.New("answer").Parse(answerTpl))
	buf := new(bytes.Buffer)
	t.Execute(buf, data)

	url, ok := data["io-messages"]
	if !ok {
		return mkMap("error", "no parameter 'io-messages'")
	}
	key, ok := data["io-apikey"]
	if !ok {
		return mkMap("error", "no parameter 'io-apikey'")
	}

	entry := Entrypoint{
		URL: url.(string),
		Key: key.(string),
	}

	message := Message{}

	s, ok := data["dest"].(string)
	if !ok {
		return mkMap("error", "no parameter 'dest'")
	}
	message.Dest = s

	s, ok = data["subject"].(string)
	if !ok {
		return mkMap("error", "no parameter 'subject'")
	}
	message.Subject = s

	s, ok = data["markdown"].(string)
	if !ok {
		return mkMap("error", "no parameter 'markdown'")
	}
	message.Markdown = s

	m, err := SendMessage(&entry, &message)
	if err != nil {
		return mkMap("error", err.Error())
	}
	return m
}
