package main

import "fmt"

func main() {

	var a = 10
	var b = 20
	var c = 30

	var d = a + b - c
	var e = a * b - c

	var f = 10
	f += 10 // f = f + 10
	f += 5 // f = f + 5

	var i = 1
	i++ // i = i + 1
	fmt.Println("Nilai i:", i)
	i++ // i = i + 1
	// i++ // i = i + 1
	fmt.Println(i)

	i-- // i = i - 1
	fmt.Println("Nilai i setelah decrement:", i)
	
 
	fmt.Println("Nilai a:", a)
	fmt.Println("Nilai b:", b)
	fmt.Println("Nilai c:", c)
	fmt.Println("Jumlah a + b - c:", d)
	fmt.Println(e)

	fmt.Println(f)
}