package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	defer catch()
	var input, output string
	var regex, _ = regexp.Compile(`\W|\d`)

	fmt.Print("Masukan input : ")
	fmt.Scan(&input)

	output = regex.ReplaceAllString(input, "")
	if  valid, err := validate(output); valid == false{
		panic(err.Error())
	}
	output = reverseString(output)
	for i, v := range output {
		if i == 0 {
			v = unicode.ToUpper(v)
		}
		if i == len(output)-1 {
			v = unicode.ToLower(v)
		}
		fmt.Printf("%c", v)
	}
}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func validate(data string) (bool, error) {
	if strings.TrimSpace(data) == "" {
		return false, errors.New("Harus terdapat string")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error(\"", r, "\")")
	}
}
