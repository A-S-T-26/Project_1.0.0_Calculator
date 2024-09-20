package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operators = [4]string{"+", "-", "/", "*"}
var convertedRomeNumbers = [14]int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var RomeToArabNumbers = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

func checkOperators(operator string) bool {
	if len(operator) == 1 {
		for _, symb := range operators {
			if symb == operator {
				return true
			}
		}
	}
	return false
}

func defineRomeNumbers(number string) bool {
	for key, value := range RomeToArabNumbers {
		if (key == number) && (value < 11) {
			return true
		}
	}
	return false
}

func getResRome(answer int) string {
	var RomeNumb string
	if answer < 0 {
		panic("В Римской системе нет отрицательных чисел.")
	} else if answer == 0 {
		panic("В Римской системе нет нуля.")
	} else {
		for _, convNumb := range convertedRomeNumbers {
			for i := convNumb; i <= answer; {
				for key, value := range RomeToArabNumbers {
					if value == convNumb {
						RomeNumb += key
						answer -= convNumb
					}
				}
			}
		}
	}
	return RomeNumb
}

func getRes(number1, number2 int, operator string) int {
	var res int
	switch operator {
	case "+":
		res = number1 + number2
	case "-":
		res = number1 - number2
	case "*":
		res = number1 * number2
	case "/":
		res = number1 / number2

	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите математический пример: ")
	example, _ := reader.ReadString('\n')
	example = strings.TrimSpace(example)
	exampleParts := strings.Split(example, " ")
	if len(exampleParts) != 3 {
		panic("Формат ввода не удовлетворяет заданию — два операнда и один оператор, разделенный единичными пробелами (+, -, /, *). Пример корректной записи: 1 + 2")
	}
	if !checkOperators(exampleParts[1]) {
		panic("Формат ввода не удовлетворяет заданию — два операнда и один оператор, разделенный единичными пробелами (+, -, /, *). Пример корректной записи: 1 + 2")
	} else if (defineRomeNumbers(exampleParts[0]) == true) && (defineRomeNumbers(exampleParts[2]) == true) {
		answer := getRes(RomeToArabNumbers[exampleParts[0]], RomeToArabNumbers[exampleParts[2]], exampleParts[1])
		fmt.Print(getResRome(answer))
	} else {
		number1, _ := strconv.Atoi(exampleParts[0])
		operator := exampleParts[1]
		number2, _ := strconv.Atoi(exampleParts[2])
		if (number1 > 0 && number1 < 11) && (number2 > 0 && number2 < 11) {
			fmt.Print(getRes(number1, number2, operator))
		} else {
			panic("Калькулятор умеет работать только с арабскими (от 1 до 10 включительно) или римскими цифрами одновременно(от I до X включительно)")
		}
	}
}
