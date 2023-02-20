package test

import (
	"fmt"
	"testing"
)

type Person struct {
	Id   int
	name string
	age  int
}

type SayHelloWorld interface {
	HelloWorld() string
}

func (person *Person) HelloWorld() string {
	return person.name + " hello World"
}

func TestInterface(t *testing.T) {
	person := Person{1, "majiang", 18}
	t.Log(person.HelloWorld())

	// Person 为 SayHelloWorld 的子类
	var person1 SayHelloWorld
	person1 = &Person{1, "majiang", 18}
	t.Log(person1.HelloWorld())
}

func TestStruct(t *testing.T) {
	person := Person{1, "majiang", 18}
	fmt.Println("person", person)

	personPoint := &person
	fmt.Println("personPoint", personPoint)
	fmt.Println("person.name", person.name)
	fmt.Println("personPoint.name", personPoint.name)
	fmt.Println("&person.name", &person.name)
	fmt.Println("&personPoint.name", &personPoint.name)
}
