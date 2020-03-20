package wskide

import "fmt"

func ExampleWhiskDockerRunOk() {
	//*DryRunFlag = false
	DryRunPush("", "1.2.3.4", "1234566789", "")
	fmt.Println(whiskDockerRun())
	// Output:
	// docker pull actionloop/iosdk:latest
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:latest
	// docker exec openwhisk waitready
}

func ExampleWhiskDockerRunKo() {
	//*DryRunFlag = false
	DryRunPush("cannot pull")
	fmt.Println(1, whiskDockerRun())
	DryRunPush("", "Error: cannot find ide")
	fmt.Println(2, whiskDockerRun())

	DryRunPush("", "1.2.3.4", "!cannot start")
	fmt.Println(3, whiskDockerRun())
	DryRunPush("", "1.2.3.4", "1234", "!no wait")
	fmt.Println(4, whiskDockerRun())
	// Output:
	// docker pull actionloop/iosdk:latest
	// 1 cannot pull actionloop/iosdk:latest
	// docker pull actionloop/iosdk:latest
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// 2 cannot locate redis
	// docker pull actionloop/iosdk:latest
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:latest
	// 3 cannot start server: cannot start
	// docker pull actionloop/iosdk:latest
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:latest
	// docker exec openwhisk waitready
	// 4 server readyness error: !no wait
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
