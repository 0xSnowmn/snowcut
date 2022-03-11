package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	urls := getUrls()
	for url := range urls {
		reversed_string := Reverse(url)
		cut_reversed_string := strings.Split(reversed_string, ".")
		final := Reverse(strings.Join(cut_reversed_string[2:], "."))
		fmt.Println(final)
	}
}

// Reverse String Function
func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

// Read urls from std input in case user do 'cat file'
func getUrls() <-chan string {
	// create urls channel
	urls := make(chan string)
	// scan all std
	scan := bufio.NewScanner(os.Stdin)
	go func() {
		defer close(urls)
		for scan.Scan() {
			// send every line to urls channel
			urls <- scan.Text()
		}
	}()
	return urls
}
