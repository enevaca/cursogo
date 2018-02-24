package main

import (
  "fmt"
  "strings"
  "github.com/enevaca/gocurso/flow"
  "github.com/enevaca/gocurso/name"
  "github.com/enevaca/gocurso/numbers"
  "github.com/enevaca/gocurso/structs"
)

const helloWorld string = "Hola %s %s,  bienvenido al fascinante mundo de Go\n"
const testConst = "Test"

func main() {
  lastname := "<Modificar con mi apellido>"
  firstname := name.GetName()
  number := numbers.Sum(50, 50) 
  a, b, c := numbers.GetVariables()
  f32, f64 := numbers.GetFloat()
  stringUTF8 := name.GetUnicode()
  
  fmt.Printf(helloWorld, firstname, lastname)
  fmt.Println("Hola mundo")
  fmt.Println(number, a, b, c)
  fmt.Println(f32, f64)
  fmt.Println("Cadena con UTF8: ", stringUTF8)
  fmt.Println(string("hola"[0]))
  fmt.Println("Cantidad de caracteres de hola: ", len("hola"))
  structs.GetArray()
  structs.GetSlice()
  flow.IfTest()
  forTest()
  string2()
  flow.SwitchTest()
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
