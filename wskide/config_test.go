package wskide

/*

	- Input()
	- ConfigLoad()
	- ConfigSave()

	- InitRepo()

	iosdk start

		ConfigLoad()
		if empty: please run init and exit
		start ide using home

	iosdk init <directory>
			ConfigLoad()
			err, dir = InitRepo(dir)
			Config.AppDir = dir
			Configure()
			ConfigSave()

			if directory not empty:
			then use it
			else: ask for repository to use
				clone it
			fi


*/

func ExampleLoadConfig() {
	// Empty load config
	//_ := ConfigLoad()
	// should ask for missing values
	//fmt.Println(cfg, err)
	// Output:
	// -
}

func ExampleConfigure() {
	Configure()
	// Output:
	// -
}
