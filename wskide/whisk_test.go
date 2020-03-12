package wskide

import "fmt"

func ExampleWhiskDockerRunOk() {
	//*DryRunFlag = false
	DryRunPush("", "1234566789", "")
	fmt.Println(whiskDockerRun())
	// Output:
	// docker pull actionloop/iosdk:latest
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:latest
	// docker exec openwhisk waitready
}

func ExampleWhiskDockerRunKo() {
	//*DryRunFlag = false
	DryRunPush("cannot pull")
	fmt.Println(whiskDockerRun())
	DryRunPush("", "!cannot start")
	fmt.Println(whiskDockerRun())
	DryRunPush("", "1234", "!no wait")
	fmt.Println(whiskDockerRun())
	// Output:
	// docker pull actionloop/iosdk:latest
	// cannot pull actionloop/iosdk:latest
	// docker pull actionloop/iosdk:latest
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:latest
	// cannot start server: cannot start
	// docker pull actionloop/iosdk:latest
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:latest
	// docker exec openwhisk waitready
	// server readyness error: !no wait
}

func ExampleWhiskDockerRm() {
	// *DryRunFlag = false
	fmt.Println(WhiskDestroy())
	// Output:
	// Destroying Whisk...
	// docker exec openwhisk stop
	//
	// <nil>
}
