package main

import (
	"fmt"
)

var entry = []Entrypoint{
	Entrypoint{
		URL: "http://localhost:3280/api/v1/web/guest/util/messages",
		Key: "123456",
	},
	Entrypoint{
		URL: "https://api.cd.italia.it/api/v1/messages",
		Key: "483b7b1f3a974b45b5c44a43538c9255",
	},
}

func ExampleSendMessage() {
	n := 0
	msg := Message{}
	res, err := SendMessage(&entry[n], &msg)
	_, ok := res["detail"]
	fmt.Println(1, ok, err)

	msg.Subject = "Welcome new user !"
	res, err = SendMessage(&entry[n], &msg)
	_, ok = res["detail"]
	fmt.Println(2, ok, err)

	msg.Markdown = `# This is a markdown header\nto show how easily markdown can be converted to **HTML**\n\nRemember: this has to be a long text.`
	res, err = SendMessage(&entry[n], &msg)
	_, ok = res["detail"]
	fmt.Println(3, ok, err)

	msg.Dest = "ISPXNB32R82Y766D"
	res, err = SendMessage(&entry[n], &msg)
	_, ok = res["id"]
	fmt.Println(4, ok, err)
	// Output:
	// 1 true <nil>
	// 2 true <nil>
	// 3 true <nil>
	// 4 true <nil>
}
