package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// scanning()
	// createNewTxt()
	// readNewTxt()
	// loggingErrorMsg()
	// topics.Recovering()
	// errorMsg()
	errorMsgFormatted()
}

func scanning() {
	fmt.Println("\n\nscanning:")

	var answer1, answer2 string

	fmt.Print("\nName: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("\n", answer1, " loves to eat ", answer2)
}

func createNewTxt() {
	fmt.Println("\n\nCreating new-text.txt using createNewTxt()...")

	f, err := os.Create("10-error-handling/new-text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	r := strings.NewReader("From createNewTxt() in 10-error-handling main.go file.")
	io.Copy(f, r)
}

func readNewTxt() {
	fmt.Println("\n\nreadNewTxt:")

	f, err := os.Open("10-error-handling/new-text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

func loggingErrorMsg() {
	fmt.Println("\n\nloggingErrorMsg into error-logs.txt:")

	f, err := os.Create("10-error-handling/error-logs.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("does-not-exist.txt")
	if err != nil {
		// fmt.Println("error:", err)
		// panic(err)

		// Check error-logs.txt
		// log.Print("error:\n", err)
		log.Fatal("error:\n", err)
	}

	defer f2.Close()
}

func errorMsg() {
	fmt.Println("\n\nerrorMsg:")

	sqrt := func(f float64) (float64, error) {
		if f < 0 {
			return 0, errors.New("error: square root of negative number")
		}

		return 42, nil
	}

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func errorMsgFormatted() {
	sqrt := func(f float64) (float64, error) {
		if f < 0 {
			return 0, fmt.Errorf("square root of negative number: %v", f)
		}

		return 42, nil
	}

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
