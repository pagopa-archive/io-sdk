package wskide

// Start openwhisk-ide
func Start(dir string) error {
	err := Preflight(dir)
	if err != nil {
		return err
	}
	err = WhiskDeploy()
	if err != nil {
		return err
	}
	err = IdeDeploy(dir)
	if err != nil {
		return err
	}
	return nil
}

// Stop openwhisk-ide
func Stop() error {
	IdeDestroy()
	WhiskDestroy()
	return nil
}
