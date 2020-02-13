package wskide

import (
	"fmt"
)

func ExampleStart() {
	*DryRunFlag = true
	DryRunPush(MinDockerVersion, "", "172.17.0.2")
	fmt.Println(Start(""))
	// Output:
	// docker version --format {{.Server.Version}}
	// Deploying Whisk...
	// docker run -d -p 3232:3232 -p 3233:3233 --rm --name openwhisk --hostname openwhisk -v /var/run/docker.sock:/var/run/docker.sock -e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 openwhisk/standalone:nightly
	// Done.
	// Deploying IDE...
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} openwhisk
	// docker run -d -p 3000:3000 --rm --name ide-js -v /var/run/docker.sock:/var/run/docker.sock --add-host=openwhisk:172.17.0.2  actionloop/ide-js
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
	// Done.
	// Destroying Whisk...
	// docker exec openwhisk stop
	//
	// Done.
	// <nil>
}
