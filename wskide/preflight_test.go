package wskide

import (
	"fmt"
	"os"
)

func ExampleEnsureDockerVersion() {
	// *DryRunFlag = false
	DryRunPush("19.03.5", "10.03.5", MinDockerVersion)
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
}

func ExampleInHomePath() {
	// *DryRunFlag = false
	os.Setenv("HOME", "/tmp")
	fmt.Println(preflightInHomePath("/tmp/openwhisk-ide"))
	fmt.Println(preflightInHomePath("/var/run"))
	fmt.Println(preflightInHomePath(""))
	// Output:
	// <nil>
	// work directory /var/run should be under your home directory; this is required to be accessible by Docker
	// <nil>
}
