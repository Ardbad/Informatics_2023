package lab5

import (
	"fmt"
)

type Dog struct {
	age   int
	name  string
	breed string
}

func (d Dog) Getage() int {
	return d.age
}

func (d Dog) Getname() string {
	return d.name
}

func (d Dog) Getbreed() string {
	return d.breed
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d *Dog) SetAge(age int) {
	d.age = age
}

func Thedog(name string, age int, breed string) (*Dog, error) {
	d := &Dog{
		name: name,
		age:  age,
	}
	var err = d.Setage(age)
	return d, err
}

func (d *Dog) Setage(age int) error {
	if age > 0 && age < 20 {
		d.age = age
		return nil
	}

	return fmt.Errorf(d.Getname(), "age is incorrect", d.Getage())
}
