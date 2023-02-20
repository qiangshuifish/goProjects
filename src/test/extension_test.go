package test

import (
	"fmt"
	"testing"
)

type Pet struct {
	name string
}

type Dog struct {
	p *Pet
}

type Cat struct {
	Pet
}

type Speak interface {
	Speak() string
}

func (pet *Pet) Speak() string {
	return "Pet Speak " + pet.name
}

func (cat *Cat) Speak() string {
	return "Cat Speak " + cat.name
}

func TestExtension(t *testing.T) {
	pet := Pet{"pet"}
	t.Log("pet.Speak()", pet.Speak())

	dog := Dog{&Pet{"dog"}}
	t.Log("dog.p.Speak()", dog.p.Speak())

	cat := Cat{Pet{"麻酱"}}
	t.Log("cat.Speak()", cat.Speak())

	t.Log("=================================")
	PrintSpeak(&pet)
	PrintSpeak(dog.p)
	PrintSpeak(&cat)
}

func PrintSpeak(speak Speak) {
	fmt.Println(speak.Speak())
}
