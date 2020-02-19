package wskide

func ExamplePlayground() {
	*DryRunFlag = true
	ideDockerRun("")

	// Output:
	// docker pull actionloop/ide-js:latest
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} openwhisk
	// docker run -d -p 3000:3000 --rm --name ide-js --add-host=openwhisk: actionloop/ide-js:latest
}
