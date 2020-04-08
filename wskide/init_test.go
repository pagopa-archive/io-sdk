package wskide

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ExampleInit2() {
	run("rm -Rf /tmp/io-sdk")
	os.Setenv("HOME", "/tmp/io-sdk")
	DryRunPush("/tmp/io-sdk/javascript", "javascript")
	dir, err := Init("", "", ioutil.Discard)
	fmt.Println(1, dir, err)
	DryRunPush("")
	dir, err = Init("", "", ioutil.Discard)
	fmt.Println(2, err)
	DryRunPush("/tmp/io-sdk/javascript", "")
	dir, err = Init("", "", ioutil.Discard)
	fmt.Println(3, err, dir)
	// Output:
	// Select one of the available templates for importers, or provide your own.
	// The javascript template is for Excel import.
	// The java template is for SQL import.
	// The python template is for REST import.
	// The github template requires a github repo (user/path).
	// Preparing work directory /tmp/io-sdk/javascript for https://github.com/pagopa/io-sdk-javascript
	// Done.
	// 1 /tmp/io-sdk/javascript <nil>
	// 2 Directory not specified
	// Using existing work directory.
	// 3 <nil> /tmp/io-sdk/javascript
}

func ExampleInit() {
	//*DryRunFlag = false
	run("rm -Rf /tmp/io-sdk")
	os.Setenv("HOME", "/tmp/io-sdk")
	fmt.Println(0, last(Init("/var/tmp/demo", "", ioutil.Discard)))
	DryRunPush("javascript")
	fmt.Println(1, last(Init("/tmp/io-sdk/demo", "", ioutil.Discard)))
	fmt.Print(run("ls /tmp/io-sdk"))
	fmt.Println(2, last(Init("/tmp/io-sdk/demo", "", ioutil.Discard)))
	run("rm -Rf /tmp/io-sdk")
	fmt.Println(3, last(Init("/tmp/io-sdk/clone", "pagopa/io-sdk-javascript", ioutil.Discard)))
	fmt.Print(run("ls /tmp/io-sdk"))
	fmt.Println(4, last(Init("/tmp/io-sdk/donotexist", "pagopa/do-not-exist", ioutil.Discard)))
	// Output:
	// 0 work directory /var/tmp/demo should be under your home directory; this is required to be accessible by Docker
	// Select one of the available templates for importers, or provide your own.
	// The javascript template is for Excel import.
	// The java template is for SQL import.
	// The python template is for REST import.
	// The github template requires a github repo (user/path).
	// Preparing work directory /tmp/io-sdk/demo for https://github.com/pagopa/io-sdk-javascript
	// Done.
	// 1 <nil>
	// demo
	// Using existing work directory.
	// 2 <nil>
	// Preparing work directory /tmp/io-sdk/clone for https://github.com/pagopa/io-sdk-javascript
	// Done.
	// 3 <nil>
	// clone
	// Preparing work directory /tmp/io-sdk/donotexist for https://github.com/pagopa/do-not-exist
	// 4 authentication required
}

// Preparing work directory /tmp/io-sdk/demo for javascript
// 1 <nil>
// Preparing work directory /tmp/io-sdk/non-exists-directory for non-exists-language
// 2 authentication required
// Preparing work directory /tmp/io-sdk/demo for javascript
// 3 repository already exists
