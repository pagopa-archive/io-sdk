package wskide

import "fmt"

func ExampleInit() {
	//*DryRunFlag = false
	DryRunPush("repository already exists")
	fmt.Println(Init("project", "javascript"))
	// Output:
	// repository already exists
}
