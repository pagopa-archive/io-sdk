package wskide

import (
	"fmt"
)

var (
	whiskDeployCmd  = debugCmd.Command("whisk-deploy", "Create Whisk deployment").Hidden()
	whiskDestroyCmd = debugCmd.Command("whisk-destroy", "Destroy Whisk deployment").Hidden()
)

func whiskParse(cmd string) bool {
	switch cmd {
	case whiskDeployCmd.FullCommand():
		WhiskDeploy()
		return true
	case whiskDestroyCmd.FullCommand():
		WhiskDestroy()
		return true
	}
	return false
}

// WhiskDeploy deploys openwhisk standalone
func WhiskDeploy() error {
	fmt.Println("Deploying Whisk...")
	whiskDockerRun()
	fmt.Println("Done.")
	return nil
}

// WhiskDestroy destroys openwhisk standalone
func WhiskDestroy() error {
	fmt.Println("Destroying Whisk...")
	fmt.Println(Sys("docker exec openwhisk stop"))
	fmt.Println("Done.")
	return nil
}

func whiskDockerRun() string {
	fmt.Printf("docker run -d -p 3232:3232 -p 3233:3233 --rm --name openwhisk --hostname openwhisk -v /var/run/docker.sock:/var/run/docker.sock -e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 openwhisk/standalone:nightly")
	return Sys(`docker run -d -p 3232:3232
-p 3233:3233 --rm --name openwhisk
--hostname openwhisk -v /var/run/docker.sock:/var/run/docker.sock
-e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 openwhisk/standalone:nightly`)
}
