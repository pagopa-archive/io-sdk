package wskide

import (
	"fmt"
	"os"
)

func ExampleConfigLoad() {
	run("rm -Rvf /tmp/iosdk-test ; mkdir /tmp/iosdk-test")
	os.Setenv("HOME", "/tmp/iosdk-test")
	fmt.Println(ConfigLoad())
	DryRunPush("123456")
	Configure("/tmp/iosdk-test/javascript")
	fmt.Println(ConfigLoad())
	fmt.Print(run("ls -a /tmp/iosdk-test/.io*"))
	fmt.Println(Config.IoAPIKey)
	fmt.Println(len(Config.WhiskAPIKey))
	// Output:
	// stat /tmp/iosdk-test/.iosdk: no such file or directory
	// <nil>
	// /tmp/iosdk-test/.iosdk
	// 123456
	// 101

}

func ExampleConfigure() {
	//Configure()
	// Output:
	// -
}
