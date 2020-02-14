package wskide

import "fmt"

func ExampleInit() {
	*DryRunFlag = false
	//DryRunPush()
	Sys("rm -Rf /tmp/io-sdk/demo")
	fmt.Println(Init("/tmp/io-sdk/demo", "javascript"))
	fmt.Println(Init("/tmp/io-sdk/non-exists-directory", "non-exists-language"))
	fmt.Println(Init("/tmp/io-sdk/demo", "javascript"))
	// Output:
	// <nil>
	// authentication required
	// repository already exists
}
