package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		interpreter := Interpreter{scanner.Text(), 0, Token{0, "", ""}}
		result := interpreter.expr()
		fmt.Println(result)
	}
}

type GeneralError struct {
	message string
}

func (e *GeneralError) Error() string {
	return fmt.Sprintf(e.message)
}

type Token struct {
	value    int
	operator string
	style    string
}

type Interpreter struct {
	text         string
	pos          int
	currentToken Token
}

func (i *Interpreter) getNextToken() (Token, error) {
	// lexical analyzer
	text := i.text

	// return EOF token because we're at the end of the text
	if i.pos > len(text)-1 {
		return Token{0, "", EOF}, nil
	}

	// get character at current position
	currentChar := string(text[i.pos])

	for currentChar == " " {
		i.pos++
		currentChar = string(text[i.pos])
	}

	// if the character is a digit convert it to an int
	if isNumeric(currentChar) == true {
		token := Token{i.getNumber(), "", INTEGER}
		return token, nil
	}

	if currentChar == "+" {
		token := Token{0, currentChar, PLUS}
		i.pos += 1
		return token, nil
	}

	if currentChar == "-" {
		token := Token{0, currentChar, MINUS}
		i.pos += 1
		return token, nil
	}

	return Token{0, "", ""}, &GeneralError{"failed to get next token"}
}

func (i *Interpreter) getNumber() int {
	currentChar := string(i.text[i.pos])
	isNumber := isNumeric(currentChar)
	totalChars := ""
	for isNumber == true {
		totalChars += currentChar
		i.pos += 1
		if i.pos > len(i.text)-1 {
			isNumber = false
		} else {
			currentChar = string(i.text[i.pos])
			isNumber = isNumeric(currentChar)
		}
	}
	return getNumber(totalChars)
}

func (i *Interpreter) eat(tokenStyle string) (Token, error) {
	if i.currentToken.style == tokenStyle {
		token, err := i.getNextToken()
		if err == nil {
			i.currentToken = token
			return i.currentToken, nil
		} else {
			return Token{0, "", ""}, &GeneralError{"failed to eat token"}
		}
	} else {
		return Token{0, "", ""}, &GeneralError{"failed to eat token"}
	}
}

func (i *Interpreter) expr() int {
	// set the current token to the first token in the input
	token, err := i.getNextToken()
	if err == nil {
		i.currentToken = token
	}

	left := i.currentToken

	i.eat(INTEGER)

	_, operr := i.eat(PLUS)
	if operr != nil {
		i.eat(MINUS)
		fmt.Println("Minus")
		fmt.Println(operr)
	} else {
		fmt.Println("Plus")
	}

	right := i.currentToken

	i.eat(INTEGER)

	fmt.Println(left, right)

	var result int

	if operr == nil {
		result = left.value + right.value
	} else {
		result = left.value - right.value
	}

	return result
}
