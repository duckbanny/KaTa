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
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}

	if len(userInputSlice) > 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
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

	a, aIsRoman := romanNum[userInputSlice[0]]
	b, bIsRoman := romanNum[userInputSlice[2]]
	resultToRoman := false

	if aIsRoman && bIsRoman {
		resultToRoman = true
	} else {
		var aIsInt, bIsInt error
		a, aIsInt = strconv.Atoi(userInputSlice[0])
		b, bIsInt = strconv.Atoi(userInputSlice[2])

		if aIsInt != nil || bIsInt != nil {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
			return
		}
		if !((a >= 1 && a <= 10) && (b >= 1 && b <= 10)) {
			println("Вывод ошибки, числа должны быть от 1 до 10")
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
		if result < 1 {
			println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			return
		}
		romanResult := intToRoman(result)
		fmt.Println(romanResult)
	} else {
		fmt.Println(result)
	}
}
