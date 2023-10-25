package stdin_wrapper

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
	"regexp"
)

type StdinReader struct {
	value			string
	helpText		string
	regexPattern	string
	notValidStr		string
	typeName		string
}


func NewStdinReader(helpText string, regexPattern string, notValidStr string) StdinReader {	
	return StdinReader{
		helpText: helpText,
		regexPattern: regexPattern,
		notValidStr: notValidStr,
	}
}

func (std *StdinReader) ReadString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(std.helpText)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	var inputStr string = strings.TrimSpace(input)
	std.value = inputStr
	return inputStr, nil
}

func (std *StdinReader) IsValidByRegex() bool{
	if std.regexPattern == "" {
		return true
	}
	match, _ := regexp.MatchString(std.regexPattern, std.value)
	if !match {
		fmt.Printf("Некорретное значение, допустимы %v\n", std.notValidStr)
	}
	return match
}


