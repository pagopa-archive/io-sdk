package wskide

import (
	"fmt"
)

func ExampleStart() {
	//*DryRunFlag = true
	DryRunPush(MinDockerVersion, "", "123", "", "456", "", "", "172.17.0.2")
	fmt.Println(Start(""))
	// Output:
	// docker version --format {{.Server.Version}}
	// Deploying Redis...
	// docker pull library/redis:5
	// docker run -d --rm --name redis --hostname redis library/redis:5
	//
	// Deploying Whisk...
	// docker pull actionloop/iosdk:latest
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:latest
	// docker exec openwhisk waitready
	//
	// Deploying IDE...
	// docker pull actionloop/ide-js:latest
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} openwhisk
	// docker run -d -p 3000:3000 --rm --name ide-js --add-host=openwhisk:172.17.0.2 actionloop/ide-js:latest
	// <nil>
}

func ExampleStop() {
	*DryRunFlag = true
	DryRunPush()
	fmt.Println(Stop())
	// Output:
	// Destroying IDE...
	// docker kill ide-js
	//
	// Destroying Whisk...
	// docker exec openwhisk stop
	//
	// Destroying Redis...
	// docker stop redis
	//
	// <nil>
}
