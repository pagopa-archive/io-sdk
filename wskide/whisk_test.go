package wskide

import "fmt"

func ExampleWhiskDockerRunOk() {
	//*DryRunFlag = false
	DryRunPush("", "1.2.3.4", "1234566789", "")
	fmt.Println(whiskDockerRun())
	// Output:
	// docker pull actionloop/iosdk:testing
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=40c41995-8221-4ec6-9e79-d82cbd7e26e5:Kz0ZjTJXJ2A5fWWcQyIdn75gIrwmDv78x0Z259CDNH32kfYsPiQWr72NiUt6LctQ -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:testing
	// docker exec openwhisk wsk property set apihost http://localhost:3233 --apihost http://localhost:3233 auth 40c41995-8221-4ec6-9e79-d82cbd7e26e5:Kz0ZjTJXJ2A5fWWcQyIdn75gIrwmDv78x0Z259CDNH32kfYsPiQWr72NiUt6LctQ --auth 40c41995-8221-4ec6-9e79-d82cbd7e26e5:Kz0ZjTJXJ2A5fWWcQyIdn75gIrwmDv78x0Z259CDNH32kfYsPiQWr72NiUt6LctQ
	// docker exec openwhisk waitready
	// docker cp /home/nimmichele/.iosdk openwhisk:/tmp/.iosdk
	// docker exec openwhisk wsk package update iosdk -P /tmp/.iosdk
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
	// docker pull actionloop/iosdk:testing
	// 1 cannot pull actionloop/iosdk:testing
	// docker pull actionloop/iosdk:testing
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// 2 cannot locate redis
	// docker pull actionloop/iosdk:testing
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=40c41995-8221-4ec6-9e79-d82cbd7e26e5:Kz0ZjTJXJ2A5fWWcQyIdn75gIrwmDv78x0Z259CDNH32kfYsPiQWr72NiUt6LctQ -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:testing
	// 3 cannot start server: cannot start
	// docker pull actionloop/iosdk:testing
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} redis
	// docker run -d -p 3280:3280 --rm --name openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=40c41995-8221-4ec6-9e79-d82cbd7e26e5:Kz0ZjTJXJ2A5fWWcQyIdn75gIrwmDv78x0Z259CDNH32kfYsPiQWr72NiUt6LctQ -v //var/run/docker.sock:/var/run/docker.sock actionloop/iosdk:testing
	// docker exec openwhisk wsk property set apihost http://localhost:3233 --apihost http://localhost:3233 auth 40c41995-8221-4ec6-9e79-d82cbd7e26e5:Kz0ZjTJXJ2A5fWWcQyIdn75gIrwmDv78x0Z259CDNH32kfYsPiQWr72NiUt6LctQ --auth 40c41995-8221-4ec6-9e79-d82cbd7e26e5:Kz0ZjTJXJ2A5fWWcQyIdn75gIrwmDv78x0Z259CDNH32kfYsPiQWr72NiUt6LctQ
	// 4 cannot update properties: !no wait
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
