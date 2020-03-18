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

	// MessageSubject := fmt.Sprintf("%v", data["MessageSubject"])
	// MessageSubject = "Welcome new user!"
	// Message := fmt.Sprintf("%v", data["Message"])
	// CodFiscDest := fmt.Sprintf("%v", data["CodFiscDest"])
	// APIKeyIO := fmt.Sprintf("%v", data["ApiKeyIO"])
	// SendMessage(MessageSubject, Message, CodFiscDest, APIKeyIO)

	// result
	res := make(map[string]interface{})
	res["body"] = buf.String()
	return res
}
