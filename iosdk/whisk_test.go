package main

import "fmt"

func ExampleWhiskDockerRunOk() {
	//*DryRunFlag = false

	DryRunPush("", "1.2.3.4", "1234566789", "")
	fmt.Println(whiskDockerRun())
	// Output:
	// docker pull iosdk/iosdk-openwhisk:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP -v //var/run/docker.sock:/var/run/docker.sock iosdk/iosdk-openwhisk:test
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
	// docker pull iosdk/iosdk-openwhisk:test
	// 1 cannot pull iosdk/iosdk-openwhisk:test
	// docker pull iosdk/iosdk-openwhisk:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// 2 cannot locate redis
	// docker pull iosdk/iosdk-openwhisk:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP -v //var/run/docker.sock:/var/run/docker.sock iosdk/iosdk-openwhisk:test
	// 3 cannot start server: cannot start
	// docker pull iosdk/iosdk-openwhisk:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP -v //var/run/docker.sock:/var/run/docker.sock iosdk/iosdk-openwhisk:test
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
