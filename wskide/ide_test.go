package wskide

import "fmt"

func ExamplePlayground() {
	*DryRunFlag = false
	ideDockerRun("")

	// Output:
	//
}

func ExampleIdeDockerRun() {
	// *DryRunFlag = false
	DryRunPush("172.17.0.2", "641792b3e0112c8fa1896b8944a846dbbab88fe5729f3d464e71475afd9e6057", "Error:", "172.17.0.2")
	fmt.Println(1, ideDockerRun("/tmp"))
	fmt.Println(2, ideDockerRun("/tmp"))
	fmt.Println(3, ideDockerRun(""))
	// Output:
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} openwhisk
	// docker run -d -p 3000:3000 --rm --name ide-js -v /var/run/docker.sock:/var/run/docker.sock -v /tmp:/home/project actionloop/ide-js --add-host 172.17.0.2
	// 1 <nil>
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} openwhisk
	// 2 Error:
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} openwhisk
	// docker run -d -p 3000:3000 --rm --name ide-js -v /var/run/docker.sock:/var/run/docker.sock  actionloop/ide-js --add-host 172.17.0.2
	// 3 <nil>

}

func ExampleIdeDockerRm() {
	// *DryRunFlag = false
	fmt.Println(IdeDestroy())
	// Output:
	// Destroying IDE...
	// docker kill ide-js
	//
	// Done.
	// <nil>
}
