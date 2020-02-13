package wskide

import (
	"fmt"
	"path/filepath"
	"strings"
)

var (
	ideDeployCmd  = debugCmd.Command("ide-deploy", "Create IDE deployment").Hidden()
	ideDestroyCmd = debugCmd.Command("ide-destroy", "Destroy IDE deployment").Hidden()
)

func ideParse(cmd string) bool {
	switch cmd {
	case ideDeployCmd.FullCommand():
		IdeDeploy(*startDirArg)
		return true
	case ideDestroyCmd.FullCommand():
		IdeDestroy()
		return true
	}
	return false
}

// IdeDeploy deploys and mounts a folder
func IdeDeploy(dir string) error {
	fmt.Println("Deploying IDE...")
	if dir != "" {
		err := preflightInHomePath(dir)
		if err != nil {
			return err
		}
	}
	return ideDockerRun(dir)
}

// IdeDestroy destroys ide
func IdeDestroy() error {
	fmt.Println("Destroying IDE...")
	fmt.Println(Sys("docker kill ide-js"))
	fmt.Println("Done.")
	return nil
}

// ideDockerRun starts the ide
// it also mounts the project folder if the directory is not empty
func ideDockerRun(dir string) error {
	var err error
	mount := ""
	if dir != "" {
		dir, err = filepath.Abs(dir)
		LogIf(err)
		if err == nil {
			mount = fmt.Sprintf("-v %s:/home/project", dir)
		}
	}
	openwhiskIP := Sys("docker inspect", "--format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}", "openwhisk")
	openwhiskIP = strings.TrimSpace(openwhiskIP)
	if strings.HasPrefix(openwhiskIP, "Error:") {
		return fmt.Errorf("%s", openwhiskIP)
	}
	command := fmt.Sprintf(`docker run -d -p 3000:3000 --rm --name ide-js
-v /var/run/docker.sock:/var/run/docker.sock --add-host=openwhisk:%s %s
actionloop/ide-js`, openwhiskIP, mount)
	//OpenWhiskDockerWait()
	Sys(command)
	return nil
}

// OpenWhiskDockerWait wait for openwhisk to be
func OpenWhiskDockerWait() error {
	fmt.Println(Sys("docker exec openwhisk waitready"))
	return nil
}
