package wskide

import "fmt"

func ExampleDockerVersion() {
	// *DryRunFlag = false
	DryRunPush("19.03.5")
	fmt.Println(dockerVersion())
	// Output:
	// docker version --format {{.Server.Version}}
	// 19.03.5
}
