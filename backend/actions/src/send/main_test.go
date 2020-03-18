package main

import "fmt"

func ExampleMain() {
	data := map[string]interface{}{
		"CodFiscDest": "123456",
		"Subject":     "Subject",
		"Message":     "HelloWord",
	}
	fmt.Printf("%v\n", Main(data))
	// Output:
	// map[body:<table>
	//  <tr><th>CodFiscDesc</th><td>123456</td></td>
	//  <tr><th>Message</th><td>HelloWord</td></td>
	// </table>
	// ]
}
