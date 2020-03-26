package main

import (
	"fmt"
	"strings"
)

func ExampleMain() {
	data := make(map[string]interface{})
	fmt.Printf("%v\n", Main(data))
	data["io-messages"] = "http://localhost:3280/api/v1/web/guest/util/messages"
	fmt.Printf("%v\n", Main(data))
	data["io-apikey"] = "1234567890"
	fmt.Printf("%v\n", Main(data))
	data["fiscal_code"] = "ABCDEF12A12A012A"
	fmt.Printf("%v\n", Main(data))
	data["subject"] = "Welcome to IO, New User"
	fmt.Printf("%v\n", Main(data))
	data["markdown"] = "This is a long text"
	fmt.Printf("%v\n", Main(data))
	// Output:
	// map[body:{"error":"no parameter 'io-messages'"}]
	// map[body:{"error":"no parameter 'io-apikey'"}]
	// map[body:{"error":"no parameter 'fiscal_code'"}]
	// map[body:{"error":"no parameter 'subject'"}]
	// map[body:{"error":"no parameter 'markdown'"}]
	// map[body:{"id":"2715563104"}]
}

func _ExampleMain2() {
	data := map[string]interface{}{
		"io-messages": "https://api.cd.italia.it/api/v1/messages",
		"io-apikey":   "483b7b1f3a974b45b5c44a43538c9255",
		"fiscal_code": "ISPXNB32R82Y766D",
		"subject":     "Welcome aboard, Luca",
		"markdown":    `This is a long text. I need at least 80 characters to make the IO API happy, in order to send a test message.`,
	}
	res := Main(data)
	//fmt.Println(res)
	fmt.Println(strings.Index(res["body"].(string), "id"))
	// Output:
	// 2
}
