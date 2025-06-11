package main

import "fmt"

func main() {

	var name1 = "Muhammad"
	var name2 = "Muhammad"

	var alfa1 = 'b'
	var alfa2 = 'a'

	var angka1 = 10
	var angka2 = 20


	var result bool = name1 == name2
	fmt.Println (result)

	var result2 bool = name1 != name2
	fmt.Println(result2)

	var result3 bool = angka1 < angka2
	fmt.Println(result3)

	var result4 bool = alfa1 > alfa2
	fmt.Println(result4)
}