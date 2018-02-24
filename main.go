package main

import (
  "fmt"
  "strings"
)

const helloWorld string = "Hola %s %s,  bienvenido al fascinante mundo de Go\n"
const testConst = "Test"

func main() {
  lastname := "<Modificar con mi apellido>"
  name := getName()
  number := sum(50, 50) 
  a, b, c := getVariables()
  f32, f64 := getFloat()
  stringUTF8 := getUnicode()
  
  fmt.Printf(helloWorld, name, lastname)
  fmt.Println("Hola mundo")
  fmt.Println(number, a, b, c)
  fmt.Println(f32, f64)
  fmt.Println("Cadena con UTF8: ", stringUTF8)
  fmt.Println(string("hola"[0]))
  fmt.Println("Cantidad de caracteres de hola: ", len("hola"))
  getArray()
  getSlice()
  ifTest()
  forTest()
  string2()
}

func getName() string {
  var name string = "Nombre test"
  name = "Sin Nombre"
  fmt.Print("Ingresa tu nombre: ")
  fmt.Scanf("%s", &name)
  return name
}

func getVariables() (int, int32, int64) {
  return 1, 2147000000, 93131313131313131
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func sum(a int, b int) int {
  return a + b	
}

func getUnicode() string {
  return "もしもし！"
}

func getArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 3}
	arr1[0] = "array"
	arr1[1] = "array2"
	fmt.Println(arr1, arr2)
}

func getSlice() {
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}

func ifTest() {
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

func forTest() {
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

func string2() {
	var text = "Hello world, Hello Platzi, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Split(text, ","))

}