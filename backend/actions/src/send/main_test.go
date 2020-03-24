package main

import "fmt"

func ExampleMain() {
	data := make(map[string]interface{})
	fmt.Printf("%v\n", Main(data))
	data["io-messages"] = "http://localhost:3280/api/v1/web/guest/util/messages"
	fmt.Printf("%v\n", Main(data))
	data["io-apikey"] = "1234567890"
	fmt.Printf("%v\n", Main(data))
	data["dest"] = "Subject"
	fmt.Printf("%v\n", Main(data))
	data["subject"] = "1234567890"
	fmt.Printf("%v\n", Main(data))
	data["markdown"] = "Hello, world."
	fmt.Printf("%v\n", Main(data))
	// Output:
	// map[error:no parameter 'io-messages']
	// map[error:no parameter 'io-apikey']
	// map[error:no parameter 'dest']
	// map[error:no parameter 'subject']
	// map[error:no parameter 'markdown']
	// map[id:879953894]
}

func _ExampleMain() {
	data := map[string]interface{}{
		"io-messages": "https://api.cd.italia.it/api/v1/messages",
		"io-apikey":   "483b7b1f3a974b45b5c44a43538c9255",
		"dest":        "ISPXNB32R82Y766D",
		"subject":     "Welcome, new user",
		"markdown":    `# This is a markdown header\nto show how easily markdown can be converted to **HTML**\n\nRemember: this has to be a long text.`,
	}
	fmt.Printf("%v\n", Main(data))
	// Output:
	// -
}
