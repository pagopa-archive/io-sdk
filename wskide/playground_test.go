package wskide

func ExamplePlayground() {
	*DryRunFlag = true
	ideDockerRun("")

	// Output:
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} openwhisk
	// docker run -d -p 3000:3000 --rm --name ide-js -v /var/run/docker.sock:/var/run/docker.sock --add-host=openwhisk: actionloop/ide-js

}
