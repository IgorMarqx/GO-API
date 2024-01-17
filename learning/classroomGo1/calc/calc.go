package calc

import "fmt"

func Operation(operation string) int {
	num1 := 5
	num2 := 2

	result := num1 - num2

	if operation == "multiplicar" {
		result = num1 * num2
	} else if operation == "somar" {
		result = num1 + num2
	} else {
		fmt.Println("Operação não reconhecida")
		return 0
	}

	fmt.Println("O resultado da", operation, "é", result)
	return result
}
