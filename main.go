package main

import (
	"fmt"
	"strings"
	//"github.com/enevaca/gocurso/structs"
	/*"github.com/enevaca/gocurso/flow"
	  "github.com/enevaca/gocurso/name"
	  "github.com/enevaca/gocurso/numbers"
	  "github.com/enevaca/gocurso/structs"
	  "github.com/enevaca/gocurso/maps"*/
)



const helloWorld string = "Hola %s %s,  bienvenido al fascinante mundo de Go\n"
const testConst = "Test"

func main() {
	/*lastname := "<Modificar con mi apellido>"
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
	  flow.SwitchTest()*/
	//fmt.Println(maps.GetMap("Esnor"))
	//structs.InterfaceTest()
	/*number, err := numbers.Sum(50, 50)

	if err != nil {
		panic(err)
	}
	fmt.Println(number)*/
	pointerTest()
}


func string2() {
	var text = "Hello world, Hello Platzi, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Split(text, ","))

}

func pointerTest() {
	a := 100
	var b *int
	b = &a
	fmt.Println("Sin modificar")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
	//*b = 10
	pointerModify(b)
	fmt.Println("Con una modificación")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
}

func pointerModify(c *int) {
	*c = 10
}