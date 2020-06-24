package main

const answerTpl = `<table>
 <tr><th>CodFiscDesc</th><td>{{.Dest}}</td></td>
 <tr><th>Subject</th><td>{{.Subject}}</td></td>
 <tr><th>Body</th><td>{{.Markdown}}</td></td>
</table>
`

//"due_date": "{{.DueDate}}"

const messageTpl = `{
	"content": {
		"subject": "{{.Subject}}",
		"markdown": "{{.Markdown}}"
	},
	"fiscal_code": "{{.FiscalCode}}"
}`
