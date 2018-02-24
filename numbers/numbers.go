package numbers

// GetVariables devuelve 3 numeros enteros
func GetVariables() (int, int32, int64) {
  return 1, 2147000000, 93131313131313131
}

// GetFloat devuelve dos numeros de punto flotante
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

// Sum suma dos numeros enteros y devuelve el resultado
func Sum(a int, b int) int {
  return a + b	
}