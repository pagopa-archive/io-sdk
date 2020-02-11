package wskide

func dockerVersion() string {
	return Sys("@docker version --format {{.Server.Version}}")
}

func dockerImages() string {
	return Sys("docker images")
}
