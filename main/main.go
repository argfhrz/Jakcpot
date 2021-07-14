package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/GeistInDerSH/clearscreen"
)

var randoms []Random

type Random struct {
	Number int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// loop := true
	// var input string
	// for loop {
	menu()
	jackpot()
	kiriJackpot()
	time.Sleep(1 * time.Second)
	tengahJackpot()
	time.Sleep(1 * time.Second)
	kananJackpot()
	price()

	// 	fmt.Println("1. Main lagi")
	// 	fmt.Println("2. Keluar")
	// 	fmt.Print("Input <1/2>: ")
	// 	fmt.Scanln(&input)

	// 	switch true {
	// 	case input == "1":
	// 		loop = true
	// 	case input == "2":
	// 		loop = false
	// 	}

	// }

}

func kiriJackpot() {
	clearscreen.ClearScreen()
	atas()
	N1 := randoms[0]

	a := format("=", 1)
	b := format(" ", 8)
	c := format(" ", 6)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			print(a, b, b, a, c)

		}
		println()
	}

	e := format(" ", 9)
	fmt.Print(e, N1.Number)
	fmt.Println("")

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			print(a, b, b, a, c)

		}
		println()
	}
	bawah()
}

func tengahJackpot() {
	clearscreen.ClearScreen()
	atas()
	N1 := randoms[0]
	N2 := randoms[1]

	a := format("=", 1)
	b := format(" ", 8)
	c := format(" ", 6)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			print(a, b, b, a, c)

		}
		println()
	}

	e := format(" ", 9)
	f := format(" ", 23)
	fmt.Print(e, N1.Number, f, N2.Number)
	fmt.Println("")

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			print(a, b, b, a, c)

		}
		println()
	}
	bawah()
}

func kananJackpot() {
	clearscreen.ClearScreen()
	atas()
	N1 := randoms[0]
	N2 := randoms[1]
	N3 := randoms[2]

	a := format("=", 1)
	b := format(" ", 8)
	c := format(" ", 6)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			print(a, b, b, a, c)

		}
		println()
	}

	e := format(" ", 9)
	f := format(" ", 23)
	fmt.Print(e, N1.Number, f, N2.Number, f, N3.Number)
	fmt.Println("")

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			print(a, b, b, a, c)

		}
		println()
	}
	bawah()
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func jackpot() ([]Random, int) {

	var number int

	for i := 0; i <= 2; i++ {
		// rand.Seed(time.Now().UnixNano())
		number := rand.Intn(10)
		newValue := Random{
			Number: number,
		}

		randoms = append(randoms, newValue)

	}

	return randoms, number

}

func price() {
	var print int
	max := len(randoms) / 3
	if len(randoms) != 0 {
		for j := 0; j < max; j++ {
			if randoms[0+(j*3)].Number == randoms[1+(j*3)].Number {
				if randoms[1+(j*3)].Number == randoms[2+(j*3)].Number {
					print = 1
				}
			}
		}
	}
	if print == 1 {
		fmt.Println("JACKPOT SELAMAT")
	}
}

func menu() (string, string) {
	tengah()
	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)
	print("Tekan enter untuk start: ")
	inputStart, _ := reader.ReadString('\n')
	print("Tekan enter untuk stop: ")
	inputStop, _ := reader.ReadString('\n')

	jackpot()
	tengahJackpot()

	return inputStart, inputStop
}

func judul() {
	fmt.Print(align(ALIGN_CENTER, "N1", " ", 21))
	fmt.Print(align(ALIGN_CENTER, "N2", " ", 26))
	fmt.Println(align(ALIGN_CENTER, "N3", " ", 21))

}

func atas() {
	judul()

	a := format("=", 18)
	b := format(" ", 6)
	for i := 0; i < 2; i++ {
		fmt.Print(a, b)
	}
	fmt.Println(a)

}

func bawah() {
	a := format("=", 18)
	b := format(" ", 6)
	for i := 0; i < 2; i++ {
		fmt.Print(a, b)
	}
	fmt.Println(a)
}

func tengah() {
	clearscreen.ClearScreen()
	atas()

	a := format("=", 1)
	b := format(" ", 8)
	c := format("=", 1)
	d := format(" ", 6)
	for i := 0; i < 9; i++ {
		for j := 0; j < 3; j++ {
			print(a, b, b, c, d)

		}
		println()
	}
	bawah()
}

func format(char string, width int) string {
	line := ""
	for i := 0; i < width; i++ {
		line += char
	}
	return line
}

type PadType int8

const (
	ALIGN_RIGHT  PadType = 1
	ALIGN_LEFT   PadType = 2
	ALIGN_CENTER PadType = 3
)

func align(padType PadType, original string, padChar string, fullLength int) string {

	lenOriginal := len(original)
	if lenOriginal < fullLength {

		padLength := fullLength - lenOriginal
		padString := format(padChar, padLength)

		if padType == ALIGN_RIGHT {
			return padString + original

		} else if padType == ALIGN_LEFT {
			return original + padString

		} else {
			halfLength := padLength / 2
			padString := format(padChar, halfLength)
			return padString + original + padString
		}

	} else {

		if padType == ALIGN_RIGHT {
			start := (lenOriginal - fullLength) - 1
			return original[start:]

		} else if padType == ALIGN_LEFT {
			return original[0:fullLength]

		} else {
			return original[0:fullLength]
		}
	}
}
