package main

import "fmt"

const secret = "abd"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

var (
	i int
	s string
	b bool
)

func main() {
	// godotenv.Load()
	// fmt.Println(os.Getenv("GO_ENV"))
	// fmt.Println(calculator.Offset)
	// fmt.Println(calculator.Sum(1, 2))
	// fmt.Println(calculator.Multiply(1, 2))

	// var i int
	// var i int = 2
	// var i = 2
	i := 1
	ui := uint16(2)

	// fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i)

	// %v = value, %T = type
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	f := 1.23456
	s := "hello"
	b := true

	fmt.Printf("f: %[1]v %[1]T\n", f)
	fmt.Printf("s: %[1]v %[1]T\n", s)
	fmt.Printf("b: %[1]v %[1]T\n", b)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v, title: %v\n", pi, title)

	x := 10
	y := 1.23
	// x、yで型が違うの計算不可
	// z := x + y
	// xの値をfloat64型に変換して計算はOK
	z := float64(x) + y
	fmt.Println(z)

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)

	i = 2
	fmt.Printf("i: %v\n", i)

	i += 1
	fmt.Printf("i: %v\n", i)

	i *= 2
	fmt.Printf("i: %v\n", i)
}

func unit16(i int) {
	panic("unimplemented")
}
