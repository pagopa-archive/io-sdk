package wskide

func ExampleHttpd() {
	*DryRunFlag = true
	Httpd(8080, "src")

	// Output:
	//
}
