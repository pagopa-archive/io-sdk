package wskide

import (
	"fmt"
)

var ()

func whiskParse(cmd string) bool {
	switch cmd {
	}
	return false
}

// WhiskDeploy deploys openwhisk standalone
func WhiskDeploy() error {
	fmt.Println("Deploying Whisk...")
	fmt.Println(whiskDockerRun())
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
	err := Run("docker pull " + OpenwhiskStandaloneImage)
	if err != nil {
		return "cannot pull " + OpenwhiskStandaloneImage
	}
	cmd := fmt.Sprintf(`docker run -d -p 3232:3232
-p 3233:3233 --rm --name openwhisk --hostname openwhisk
-v //var/run/docker.sock:/var/run/docker.sock %s`, OpenwhiskStandaloneImage)
	return Sys(cmd)
}
