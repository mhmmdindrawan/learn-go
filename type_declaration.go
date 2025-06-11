package main

import "fmt"


func main() {

	type noKTP string

	var ktpIndrawan noKTP = "1234567890"


	var contoh string = "99999999"
	var contohKTP noKTP = (noKTP)(contoh)

	fmt.Println(ktpIndrawan)
	fmt.Println(contohKTP)
}