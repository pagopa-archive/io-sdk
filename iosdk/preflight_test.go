package main

import (
	"fmt"
)

func ExampleEnsureDockerVersion() {
	*DryRunFlag = true
	DryRunPush("19.03.5", "10.03.5", MinDockerVersion, "!no docker")
	fmt.Println(preflightEnsureDockerVersion())
	fmt.Println(preflightEnsureDockerVersion())
	fmt.Println(preflightEnsureDockerVersion())
	fmt.Println(preflightEnsureDockerVersion())
	// Output:
	// docker version --format {{.Server.Version}}
	// <nil>
	// docker version --format {{.Server.Version}}
	// Installed docker version 10.3.5 is no longer supported
	// docker version --format {{.Server.Version}}
	// <nil>
	// docker version --format {{.Server.Version}}
	// no docker
}

func ExampleInHomePath() {
	// *DryRunFlag = false
	fmt.Println(preflightInHomePath("/tmp/iosdk-test/openwhisk-ide"))
	fmt.Println(preflightInHomePath("/var/run"))
	fmt.Println(preflightInHomePath(""))
	// Output:
	// <nil>
	// work directory /var/run should be below your home directory /tmp/iosdk-test;
	// this is required to be accessible by Docker
	// <nil>
}

func ExamplePreflightDockerMemory() {
	*DryRunFlag = true
	*verboseFlag = true
	DryRunPush("\nTotal Memory: 11GiB\n", "\nTotal Memory: 3GiB\n")
	fmt.Println(preflightDockerMemory())
	fmt.Println(preflightDockerMemory())
	// Output:
	// docker info
	// <nil>
	// docker info
	// IOSDK needs 4GB memory allocatable on docker
}
