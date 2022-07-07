package main

import "fmt"

type Workers interface {
	salary(int) int
}

type Full_time struct {
	basic int
}
type Contractor struct {
	basic int
}
type Freelancer struct {
	basic int
}

func (e Full_time) salary(days int) int {
	return e.basic * days
}
func (e Contractor) salary(days int) int {
	return e.basic * days
}
func (e Freelancer) salary(hours int) int {
	return e.basic * hours
}

func main() {
	a := Full_time{500}
	b := Contractor{100}
	c := Freelancer{10}

	employee := []Workers{a, b, c}
	fmt.Println(employee[0].salary(28))
	fmt.Println(employee[1].salary(28))
	fmt.Println(employee[2].salary(20))
}
