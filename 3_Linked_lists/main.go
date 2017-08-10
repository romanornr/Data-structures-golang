package main

import (
	"fmt"
)

type Student struct {
	age    int
	weight int
	name   string
	Next   *Student
}

type Teacher struct {
	age    int
	weight int
	name   string
	Next   *Student
}

func main() {

	jack := Student{12, 30, "Jack", nil}
	john := Student{14, 45, "John", &jack}

	maria := Teacher{26, 60, "Maria", &john}

	fmt.Println("The teacher is: ", maria.name)
	fmt.Println("The student behind Debrah is: ", maria.Next.name)
	fmt.Println("The student begind John is: ", john.Next.name)

}
