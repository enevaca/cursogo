package main

import "fmt"

const helloWorld string = "Hola %s %s,  bienvenido al fascinante mundo de Go\n"
const testConst = "Test"

func main() {
  lastname := "<Modificar con mi apellido>"
  name := getName()
  number := sum(50, 50) 
  a, b, c := getVariables()
  f32, f64 := getFloat()
  
  fmt.Printf(helloWorld, name, lastname)
  fmt.Println("Hola mundo")
  fmt.Println(number, a, b, c)
  fmt.Println(f32, f64)
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