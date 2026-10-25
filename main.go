package main

import "fmt"

func main() {

	// VARIABLE
	// var firstName string = "Lil"
	// var lastName string
	// lastName = "Tepu"
	// fmt.Printf("halo %s %s!\n", firstName, lastName)

	// job := "programmer"
	// job = "backend engineer"
	// fmt.Printf("my job %s!\n", job)

	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"
	// fmt.Println(first, second, third)

	// var fourt, fifth, sixth string = "empat", "lima", "enam"
	// fmt.Println(fourt, fifth, sixth)

	// seventh, eighth, ninth := "tujuh", "delapan", "sembilan"
	// fmt.Println(seventh, eighth, ninth)

	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "helo"
	// fmt.Println(one, isFriday, twoPointTwo, say)

	// _ = "belajar golang"
	// _ = "golang itu mudah"
	// nameku, _ := "Lil", "tepu"
	// fmt.Println(nameku)

	// name := new(string)
	// fmt.Println(name)
	// fmt.Println(*name)

	// x, y := 10, 20
	// x, y = y, x
	// fmt.Println(x, y)

	// i, j := 5, 1
	// i, j = i+j, i
	// fmt.Println(i, j)

	// DATA TYPE
	// var positiveNumber uint8 = 89
	// var negativeNumber = -123456545678
	// fmt.Printf("bilangan positif: %d\n", positiveNumber)
	// fmt.Printf("bilangan negative: %d\n", negativeNumber)

	// var decimalNumber = 2.62
	// fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	// fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// var exist bool = true
	// fmt.Printf("exist? %t \n", exist)

	// var message string = "Halo"
	// fmt.Printf("message: %s \n", message)

	// var pesan = `Hallo nama saya "Lil tepu".
	// Salam kenal.
	// Mari kita belajar "golang"`
	// fmt.Println(pesan)

	// const namaDepan string = "Lil"
	// fmt.Print("Halo", namaDepan, "!\n")
	// const namaBelakang = "Tepu"
	// fmt.Print("nice to meet you", namaBelakang, "!\n")

	// fmt.Println("john wick")
	// fmt.Println("john", "wick")
	// fmt.Print("john wick\n")
	// fmt.Print("john ", "wick\n")
	// fmt.Print("john", " ", "wick\n")

	// const (
	// 	square         = "Kotak"
	// 	isToday        = true
	// 	numeric  uint8 = 1
	// 	floatNum       = 2.2
	// )
	// fmt.Println(square, isToday, numeric, floatNum)
	// const (
	// 	a = "konstanta"
	// 	b
	// )
	// fmt.Println(a, b)

	// const (
	// 	today string = "senin"
	// 	sekarang
	// 	hariIni = true
	// )
	// fmt.Println(today, sekarang, hariIni)

	// const satu, dua = 1, 2
	// const three, four string = "tiga", "empat"
	// fmt.Println(satu, dua, three, four)

	// OPERATOR
	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Println(value)

	var value1 = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)
	fmt.Printf("nilai %d (%t) \n", value1, isEqual)

	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t(%t) \n", leftReverse)

	fmt.Println(min(3, 7, 1, 5))
	fmt.Println(max(3, 7, 1, 5))

	var a, b = 10, 20
	fmt.Println(min(a, b))
	fmt.Println(max(a, b))

	fmt.Println(min("banana", "apple", "cherry"))
	fmt.Println(max("banana", "apple", "cherry"))
}
