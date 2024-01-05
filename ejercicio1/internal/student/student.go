package student

import "fmt"

type student struct {
	name,
	lastname string
	dni         int
	dateOfEntry string
}

func NewStudent(name, lastname string, dni int, dateOfEntry string) student {
	return student{name: name, lastname: lastname, dni: dni, dateOfEntry: dateOfEntry}
}

func (s student) Detail() {
	fmt.Println("Name:", s.name)
	fmt.Println("lastname:", s.lastname)
	fmt.Println("dni:", s.dni)
	fmt.Println("dateOfEntry:", s.dateOfEntry)
}
