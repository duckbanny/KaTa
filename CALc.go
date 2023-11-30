package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	romanNum := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""
	for _, r := range romanNum {
		for num >= r.Value {
			result += r.Symbol
			num -= r.Value
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInputSlice := strings.Fields(userInput)

	if len(userInputSlice) < 3 {
		fmt.Println("Ошибка: строка не является математической операцией.")
		return
	}

	if len(userInputSlice) > 3 {
		fmt.Println("Ошибка: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	var (
		romanNum = map[string]int{
			"I":    1,
			"II":   2,
			"III":  3,
			"IV":   4,
			"V":    5,
			"VI":   6,
			"VII":  7,
			"VIII": 8,
			"IX":   9,
			"X":    10,
		}
	)

	a, aExist := romanNum[userInputSlice[0]]
	b, bExist := romanNum[userInputSlice[2]]
	resultToRoman := false

	if aExist && bExist {
		resultToRoman = true
	} else {
		var aErr, bErr error
		a, aErr = strconv.Atoi(userInputSlice[0])
		b, bErr = strconv.Atoi(userInputSlice[2])

		if aErr != nil || bErr != nil {
			fmt.Println("Ошибка: строка не является математической операцией или не является римской цифрой от 1 до 10.")
			return
		}
	}

	var operation = userInputSlice[1]
	var result int

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	if resultToRoman {
		romanResult := intToRoman(result)
		fmt.Println(romanResult)
	} else {
		fmt.Println(result)
	}
}
