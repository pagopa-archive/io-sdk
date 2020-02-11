package wskide

import "fmt"

/*
func ExampleWhiskParse() {
	fmt.Println(1, whiskParse("whisk"))
	fmt.Println(2, whiskParse("whisk deploy"))
	fmt.Println(3, whiskParse("whisk destroy"))
	// Output:
	// 1 false
	// Deploying Whisk
	// 2 true
	// Destroying Whisk
	// 3 true
}*/

func ExampleWhiskDockerRun() {
	//*DryRunFlag = false
	DryRunPush("991c7972fa4612c873b3804a4c334b3af66687a7f1e548a36dfdfe0c6a717cbe")
	fmt.Println(whiskDockerRun())
	// Output:
	// docker run -d -p 3232:3232 -p 3233:3233 --rm --name openwhisk --hostname openwhisk -v /var/run/docker.sock:/var/run/docker.sock -e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 openwhisk/standalone:nightly
	// 991c7972fa4612c873b3804a4c334b3af66687a7f1e548a36dfdfe0c6a717cbe
}
func ExampleWhiskDockerRm() {
	// *DryRunFlag = false
	fmt.Println(WhiskDestroy())
	// Output:
	// Destroying Whisk...
	// docker exec openwhisk stop
	//
	// Done.
	// <nil>
}
