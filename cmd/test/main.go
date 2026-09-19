package main

import "TODOS_Logger/pkg/console"

func main() {
	var num1 int
	num2 := 2
	variable := "Salut"

	println(variable)
	println(num1)
	println(num2)

	variable = "Bonjour en fait"

	println(variable)

	console.Printcln(console.RED, "%s", variable)
	console.Printcln(console.GREEN, "%s", variable)
	console.Printcln(console.BLUE, "%s", variable)
	console.Printcln(console.YELLOW, "%s", variable)

	console.Printc(console.RED, "%s", variable)
}
