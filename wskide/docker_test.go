package wskide

import "fmt"

func ExampleDockerVersion() {
	//*DryRunFlag = false
	DryRunPush("19.03.5", "!no docker")
	out, err := dockerVersion()
	fmt.Println(out, err)
	out, err = dockerVersion()
	fmt.Println(out, err)
	// Output:
	// docker version --format {{.Server.Version}}
	// 19.03.5 <nil>
	// docker version --format {{.Server.Version}}
	//  no docker
}
