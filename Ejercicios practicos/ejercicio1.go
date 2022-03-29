package main

import (
	"errors"
	"fmt"
)

func division(num1 float32, num2 float32) (float32, error) {

	if num2 == 0 {
		err := errors.New("Se ha ingresado un cero")
		return 0, err
	}

	return num1 / num2, nil
}

func main() {
	div, err := division(7, 10)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		//exit(1)
		return
	}

	fmt.Println("La division es: ", div)
}
