package main

import (
	"fmt"
	"io/ioutil"
)

func ExampleStart() {
	//*DryRunFlag = true
	fmt.Println("=== Init ===")
	_, err := Start()
	fmt.Println(err)
	DryRunPush("/tmp/iosdk-test/javascript", "javascript", "123456")
	dir, err := Init("", "", ioutil.Discard)
	fmt.Println(dir, err)
	fmt.Println(Configure(dir))
	fmt.Print(run("ls -a /tmp/iosdk-test/.io*"))
	fmt.Println("=== Start ===")
	DryRunPush("\nTotal Memory: 11GiB\n", MinDockerVersion, "", "123", "", "", "", "1.2.3.4", "", "", "", "172.17.0.2")
	_, err = Start()
	fmt.Println(err)
	// Output:
	// === Init ===
	// You need to run 'iosdk init ', first.
	// stat /tmp/iosdk-test/.iosdk.v3: no such file or directory
	// Select one of the available templates for importers, or provide your own.
	// The javascript template is for Excel import.
	// The php template is for SQL import.
	// The python template is for GraphQL import.
	// The github template requires a github repo (user/path).
	// Preparing work directory /tmp/iosdk-test/javascript for https://github.com/pagopa/io-sdk-javascript
	// Done.
	// /tmp/iosdk-test/javascript <nil>
	// Wrote /tmp/iosdk-test/.iosdk.v3
	// <nil>
	// /tmp/iosdk-test/.iosdk.v3
	// === Start ===
	// docker info
	// docker version --format {{.Server.Version}}
	// WARNING: using default OpenWhisk key
	// Deploying Redis...
	// docker pull library/redis:5
	// docker run -d -p 6379:6379 --rm --name iosdk-redis --hostname redis library/redis:5 --requirepass password
	//
	// Deploying Whisk...
	// docker pull pagopa/iosdk-openwhisk:test
	// docker pull pagopa/actionloop-python-v3.7:2020-11-16
	// docker pull pagopa/action-nodejs-v10:2020-11-16
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iosdk-redis
	// docker run -d -p 3280:3280 --rm --name iosdk-openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS_IP=1.2.3.4 -e CONTAINER_EXTRA_ENV1=__OW_REDIS_PASSWORD=password -e CONFIG_FORCE_whisk_users_guest=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP -v /var/run/docker.sock:/var/run/docker.sock pagopa/iosdk-openwhisk:test
	// docker exec iosdk-openwhisk waitready
	// Deploying IDE...
	// docker pull pagopa/iosdk-theia:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iosdk-openwhisk
	// docker run -d -p 3000:3000 --rm --name iosdk-theia -e HOME=/home/project -v /tmp/iosdk-test/javascript:/home/project --add-host=openwhisk:172.17.0.2 pagopa/iosdk-theia:test
	// <nil>
}

func ExampleStop() {
	*DryRunFlag = true
	DryRunPush("iosdk-theia\n", "iosdk-openwhisk", "iosdk-redis")
	Stop()
	// Output:
	// docker kill iosdk-theia
	// Destroying IDE: iosdk-theia
	// docker exec iosdk-openwhisk stop
	// Destroying Whisk: iosdk-openwhisk
	// docker stop iosdk-redis
	// Destroying Redis: iosdk-redis
}
