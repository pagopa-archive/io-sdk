package wskide

func OffExampleHttpd() {
	*DryRunFlag = false
	Sys("pwd")

	// Output:
	// /home/nimmichele/Work/PagoPA/io-sdk/wskide
}
