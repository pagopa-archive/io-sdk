package wskide

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ExampleInit() {
	*DryRunFlag = false
	Sys("rm -Rf /tmp/io-sdk/demo")
	os.Setenv("HOME", "/tmp")
	fmt.Println(0, Init("/var/tmp/demo", "javascript", ioutil.Discard))
	fmt.Println(1, Init("/tmp/io-sdk/demo", "javascript", ioutil.Discard))
	fmt.Println(2, Init("/tmp/io-sdk/non-exists-directory", "non-exists-language", ioutil.Discard))
	fmt.Println(3, Init("/tmp/io-sdk/demo", "javascript", ioutil.Discard))
	// Output:
	// 0 work directory /var/tmp/demo should be under your home directory; this is required to be accessible by Docker
	// Preparing work directory /tmp/io-sdk/demo for javascript
	// Done.
	// 1 <nil>
	// Preparing work directory /tmp/io-sdk/non-exists-directory for non-exists-language
	// 2 authentication required
	// Preparing work directory /tmp/io-sdk/demo for javascript
	// 3 repository already exists
}
