package main

import (
	"fmt"
)

func ExampleIdeDockerRun() {
	*DryRunFlag = true
	DryRunPush("", "172.17.0.2", "641792b3e0112c8fa1896b8944a846dbbab88fe5729f3d464e71475afd9e6057",
		"", "Error:",
		"", "172.17.0.2")
	RuntimeOS = "darwin"
	fmt.Println(1, ideDockerRun("/tmp", ""))
	fmt.Println(2, ideDockerRun("/tmp", ""))
	fmt.Println(3, ideDockerRun("", ""))
	// Output:
	// docker pull pagopa/iosdk-theia:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iosdk-openwhisk
	// docker run -d -p 3000:3000 --rm --name iosdk-theia -e HOME=/home/project -v /tmp:/home/project --add-host=openwhisk:172.17.0.2 pagopa/iosdk-theia:test
	// 1 <nil>
	// docker pull pagopa/iosdk-theia:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iosdk-openwhisk
	// 2 cannot find openwhisk
	// docker pull pagopa/iosdk-theia:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iosdk-openwhisk
	// docker run -d -p 3000:3000 --rm --name iosdk-theia -e HOME=/home/project --add-host=openwhisk:172.17.0.2 pagopa/iosdk-theia:test
	// 3 <nil>
}

func ExampleIdeDockerRm() {
	*DryRunFlag = true
	DryRunPush("iosdk-theia")
	IdeDestroy()
	// Output:
	// docker kill iosdk-theia
	// Destroying IDE: iosdk-theia
}

func ExampleFixPathDockerToolbox() {
	RuntimeOS = "linux"
	fmt.Println(fixPathDockerToolbox("a\\b", "does not matter"))
	RuntimeOS = "windows"
	fmt.Println(fixPathDockerToolbox("a\\b", "does not matter"))
	fmt.Println(fixPathDockerToolbox("a\\b", "\nOperating System: Boot2Docker etc etc"))
	fmt.Println(fixPathDockerToolbox("c:\\a\\b", "\nOperating System: Boot2Docker etc etc"))
	// Output:
	// a\b
	// a\b
	// a/b
	// //c/a/b
}
