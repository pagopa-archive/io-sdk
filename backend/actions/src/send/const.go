package main

const answerTpl = `<table>
 <tr><th>CodFiscDesc</th><td>{{.Dest}}</td></td>
 <tr><th>Subject</th><td>{{.Subject}}</td></td>
 <tr><th>Body</th><td>{{.Markdown}}</td></td>
</table>
`
const messageTpl = `{
	"time_to_live": 3600,
	"content": {
		"subject": "{{.Subject}}",
		"markdown": "{{.Markdown}}",
		"due_date": "{{.DueDate}}"
	},
	"fiscal_code": "{{.Dest}}"
}`
