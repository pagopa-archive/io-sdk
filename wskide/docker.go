package wskide

import (
	"fmt"
	"strings"
)

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
