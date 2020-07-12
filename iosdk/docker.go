package main

import (
	"fmt"
	"strings"
)

func dockerInfo() (string, error) {
	out, err := SysErr("@docker info")
	if err != nil {
		return "", fmt.Errorf("Docker is not running")
	}
	return out, nil
}

func dockerVersion() (string, error) {
	return SysErr("@docker version --format {{.Server.Version}}")
}

func dockerStatus(container string) {
	res, err := SysErr("@docker inspect --format {{.State.Status}} " + container)
	if err != nil {
		res = "not running\n"
	}
	fmt.Print(container, ": ", res)
}

//dockerIP returns the ip of a container, or nil if not found
func dockerIP(container string) *string {
	ip := Sys("docker inspect",
		"--format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}",
		container)
	ip = strings.TrimSpace(ip)
	if strings.HasPrefix(ip, "Error:") {
		return nil
	}
	return &ip
}

func dockerPull(image string) error {
	if *skipPullImages {
		fmt.Println("skipping pull")
		return nil
	}
	err := Run("docker pull " + image)
	if err != nil {
		return err
	}
	return nil
}
