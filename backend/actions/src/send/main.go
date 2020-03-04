package main

import (
	"bytes"
	"html/template"
)

const answer = `<table>
 <tr><th>CodFiscDesc</th><td>{{.CodFiscDest}}</td></td>
 <tr><th>Message</th><td>{{.Message}}</td></td>
</table>
`

// Main function for the action
func Main(data map[string]interface{}) map[string]interface{} {
	t := template.Must(template.New("answer").Parse(answer))
	buf := new(bytes.Buffer)
	t.Execute(buf, data)

	// result
	res := make(map[string]interface{})
	res["body"] = buf.String()
	return res
}
