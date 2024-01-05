package main

import (
	"ejercicio1/internal/student"
)

func main() {
	e := student.NewStudent("Tomas", "Cannatella", 43434343, "01/01/2024")

	e.Detail()
}
