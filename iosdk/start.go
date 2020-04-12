package main

import "fmt"

// Start openwhisk-ide
func Start() error {
	if err := ConfigLoad(); err != nil {
		fmt.Println("You need to run 'iosdk init ', first.")
		return err
	}
	err := Preflight(Config.AppDir)
	if err != nil {
		return err
	}
	if *useDefaultAPIKey {
		Config.WhiskAPIKey = DefaultWhiskAPIKey
		fmt.Println("WARNING: using default OpenWhisk key")
	}
	err = RedisDeploy()
	if err != nil {
		return err
	}
	err = WhiskDeploy()
	if err != nil {
		return err
	}
	err = IdeDeploy(Config.AppDir)
	if err != nil {
		return err
	}
	return nil
}

// Stop openwhisk-ide
func Stop() error {
	IdeDestroy()
	WhiskDestroy()
	RedisDestroy()
	return nil
}
