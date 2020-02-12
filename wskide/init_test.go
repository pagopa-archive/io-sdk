package wskide

func ExampleInit() {
	//*DryRunFlag = false
	DryRunPush("Enumerating objects: 5, done")
	//fmt.Println(preflightEnsureDockerVersion())
	Init()
	// Output:
	// docker version --format {{.Server.Version}}
	// <nil>
	// docker version --format {{.Server.Version}}
	// Installed docker version 10.3.5 is no longer supported
	// docker version --format {{.Server.Version}}
	// <nil>
}
