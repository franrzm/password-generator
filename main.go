package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

var defaultLenght int = 30

var (
	lowerCaseLetters = []byte("abcdefghijklmnopqrstuvwxyz")
	upperCaseLetters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers          = []byte("0123456789")
	symbols          = []byte("`~!@#$%^&*()_-+={}[]\\|:;\"'<>,.?/")
)

func generatePassword(stringLength int) string {
	randomString := make([]byte, stringLength)

	rand.Seed(time.Now().UnixNano())

	for i := range randomString {
		switch rand.Intn(4) {
		case 0:
			randomString[i] = lowerCaseLetters[rand.Intn(len(lowerCaseLetters))]
		case 1:
			randomString[i] = upperCaseLetters[rand.Intn(len(upperCaseLetters))]
		case 2:
			randomString[i] = numbers[rand.Intn(len(numbers))]
		case 3:
			randomString[i] = symbols[rand.Intn(len(symbols))]
		}
	}

	return string(randomString)
}

func main() {
	var passwordLength int
	var showPassword bool

	flag.IntVar(&passwordLength, "l", defaultLenght, "Specify the length of the password. Default length is 30.")
	flag.BoolVar(&showPassword, "s", false, "Use this param to show the password on the console. Ba default the password will be hidden.")

	flag.Parse()

	var password string = generatePassword(passwordLength)

	err := clipboard.WriteAll(password)

	if err != nil {
		panic(err)
	} else {
		println("The password has been copied to clipboard!")
		println("Go to paste it!")
	}

	if showPassword {
		println("Password: ", password)
	}
}
