package maps

// GetMap devuleve un map con llave tio string y valores tipo enteros
//func GetMap() map[string]int {
func GetMap(name string) int {
	//mapTest := make(map[string]int)
	mapTest := map[string]int{
		"Esnor":    25,
		"Roxana":   28,
		"Fernando": 16,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 100
	delete(mapTest, "llave1")
	delete(mapTest, "llave2")
	return mapTest[name]
}
