package main

import (
	"fmt"
)

func ExampleIdeDockerRun() {
	*DryRunFlag = true
	DryRunPush("", "172.17.0.2", "641792b3e0112c8fa1896b8944a846dbbab88fe5729f3d464e71475afd9e6057",
		"", "Error:",
		"", "172.17.0.2")
	fmt.Println(1, ideDockerRun("/tmp"))
	fmt.Println(2, ideDockerRun("/tmp"))
	fmt.Println(3, ideDockerRun(""))
	// Output:
	// docker pull pagopa/iosdk-theia:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iosdk-openwhisk
	// docker run -d -p 3000:3000 --rm --name iosdk-theia --add-host=openwhisk:172.17.0.2 -v /tmp:/home/project pagopa/iosdk-theia:test
	// 1 <nil>
	// docker pull pagopa/iosdk-theia:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iosdk-openwhisk
	// 2 cannot find openwhisk
	// docker pull pagopa/iosdk-theia:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iosdk-openwhisk
	// docker run -d -p 3000:3000 --rm --name iosdk-theia --add-host=openwhisk:172.17.0.2 pagopa/iosdk-theia:test
	// 3 <nil>
}

func ExampleIdeDockerRm() {
	*DryRunFlag = true
	DryRunPush("172.17.0.3")
	IdeDestroy()
	// Output:
	// Destroying IDE...
	// docker kill iosdk-theia
	// 172.17.0.3
}
