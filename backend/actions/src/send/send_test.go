package main

import "fmt"

func ExampleSendMessage() {
	fmt.Println("test disabled - need to store secrets in a safe way")
	//return
	var subject = "Welcome new user !"
	var markdown = `# This is a markdown header\n\nto show how easily markdown can be converted to **HTML**\n\nRemember: this has to be a long text.`
	// need this for the Example!!
	// markdown = strings.Replace(markdown, "\n", `\n`, -1)
	var dest = "ISPXNB32R82Y766D"
	var key = "483b7b1f3a974b45b5c44a43538c9255"
	SendMessage(subject, markdown, dest, key)
	// response Body: {"id":"01E1FHC33GSY6CE64JTGDYTYMP"}

	// Output:

}
