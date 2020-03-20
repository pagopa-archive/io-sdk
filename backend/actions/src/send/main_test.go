package main

import "fmt"

func ExampleMain() {
	data := map[string]interface{}{
		"io-apihost": "http://localhost:3280/api/v1/web/guest/util/messages",
		"io-apikey":  "1234567890",
		"dest":       "1234567890",
		"subject":    "Subject",
		"markdown":   "Hello, world.",
	}
	fmt.Printf("%v\n", Main(data))
	// Output:
	// map[id:639479525]
}
