package wskide

import "fmt"

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
