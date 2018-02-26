package structs

import "fmt"

// Platzi es una interfaz de PlatziCourse y PlatziCareer
type Platzi interface {
	Subscribe(name string)
}

func deferTest() {
	fmt.Println("La función Interface ha culminado")
}

// InterfaceTest
func InterfaceTest() {
	defer deferTest()
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"backend"}}
	platziCareer := new(PlatziCareer)
	platziCareer.Name = "GoCareer"
	platziCareer.Slug = "go"
	//platziCareer.Skills = []string{"backend", "2", "3"}
	/*platiCourse1 := new(PlatziCourse)
	platiCourse1.Name = "Go1"
	platiCourse1.Slug = "go1"
	platiCourse1.Skills = []string{"backend1"}*/
	//platziCourse.Subscribe("Esnor")
	//fmt.Println(platziCourse)
	callSubscribe(platziCareer)
	callSubscribe(platziCourse)
}

func callSubscribe(p Platzi) {
	p.Subscribe("Esnor")
}