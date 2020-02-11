package wskide

import (
	"errors"
	"fmt"
)

func ExampleLogIf() {
	fmt.Println(LogIf(nil))
	err := fmt.Errorf("generated error")
	if LogIf(err) {
		fmt.Println(LastError)
	}
	// Output:
	// false
	// generated error
}

func shouldPanic(flag bool) (err error) {
	defer Recover(&err)
	if flag {
		err := errors.New("Panic mode")
		FatalIf(err)
	}
	fmt.Println("no panic")
	return nil
}

func ExampleFatalIf() {
	fmt.Println(shouldPanic(false))
	fmt.Println(shouldPanic(true))
	// Output:
	// no panic
	// <nil>
	// Panic mode
}

func ExampleSys() {
	*DryRunFlag = false
	Sys("/bin/echo 1 2 3")
	Sys("/bin/echo 3", "4", "5")
	Sys("@sh -c", "echo foo >/tmp/foo")
	fmt.Print(Sys("cat /tmp/foo"))
	Sys("@sh -c", "echo bar >/tmp/bar")
	fmt.Print(Sys("@cat /tmp/bar"))
	*DryRunFlag = true
	// Output:
	// 1 2 3
	// 3 4 5
	// foo
	// foo
	// bar
}

func ExampleSysCd() {
	*DryRunFlag = false
	Sys("@mkdir -p /tmp/example-syscd/hello-i-am-a-dir")
	fmt.Println(SysCd("/tmp/example-syscd", "@ls"))
	*DryRunFlag = true
	// Output:
	// hello-i-am-a-dir
}

func ExampleSysSh() {
	*DryRunFlag = false
	show(SysSh("@rm -Rf /tmp/simple ; mkdir /tmp/simple ; ls -d /tmp/simple"))
	SysSh("cd /tmp/simple ; touch hello ; ls")
	*DryRunFlag = true
	// Output:
	// /tmp/simple
	//
	// hello
}

func ExampleDryRun() {
	DryRunPush("first", "second")
	fmt.Println(Sys("dummy"))
	fmt.Println(Sys("dummy", "alpha", "beta"))
	fmt.Printf("'%s'\n", Sys("dummy"))
	// Output:
	// dummy
	// first
	// dummy alpha beta
	// second
	// dummy
	// ''

}

func ExampleShowError() {
	ShowError(nil)
	ShowError(errors.New("error"))
	// Output:
	// *** ERROR: error ***
}
