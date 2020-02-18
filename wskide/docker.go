package wskide

func dockerVersion() string {
	return Sys("@docker version --format {{.Server.Version}}")
}

func dockerStatus() (err error) {
	err = Run("docker ps -a")
	if err != nil {
		return err
	}
	return nil
}
