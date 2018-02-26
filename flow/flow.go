package flow

import "fmt"

// SwitchTest solicita un numero del 1 al 10 y lo evalua
func SwitchTest() {
	var number = 0
	fmt.Print("[Switch] Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)
	switch number {
		case 1: 
		  fmt.Println("El número es 1")
		default:
		  fmt.Println("El número no es 1")
	}

	switch {
	  case number % 2 == 0:
	  	fmt.Println("El número es par")
	  default:
	  	fmt.Println("El número es impar")
	}
}

// IfTest verifica si un numero es par o impar
func IfTest() {
	var number = 0
	fmt.Print("Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)
	if  number % 2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	if number2 := 3; number2 == 3 {
		fmt.Println("Entro al condicional")
	} 
}


// ForTest
func ForTest() {
	for i := 0; i < 5; i++ {
		fmt.Println("[FOR] ", i)
	}

	c := 100
	for c > 0 {
		c -= 10
		fmt.Println("[FOR] Solo con una condicion de c > 0, ", c)
	}

	s := 1000
	for {
		s -= 1
		if s == 0 {
			fmt.Println("Termina el for 'infinito'")
			break
		}
	}
}
