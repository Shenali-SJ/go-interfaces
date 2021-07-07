package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type permanent struct {
	id int
	basic int
	bonus int
}

type contract struct {
	id int
	basic int
}

type freelancer struct {
	id int
	hours int
	rate int
}

func main() {
	p1 := permanent{
		id:    1,
		basic: 10000,
		bonus: 2000,
	}

	p2 := permanent{
		id:    2,
		basic: 12000,
		bonus: 3000,
	}

	c1 := contract{
		id:    3,
		basic: 8000,
	}

	f1 := freelancer{
		id:    4,
		hours: 20,
		rate:  87,
	}

	f2 := freelancer{
		id:    5,
		hours: 10,
		rate:  56,
	}

	employees := []SalaryCalculator{p1, p2, c1, f1, f2}
	totalExpenses(employees)
}

func (p permanent) CalculateSalary() int {
	return p.basic + p.bonus
}

func (c contract) CalculateSalary() int {
	return c.basic
}

func (f freelancer) CalculateSalary() int {
	return f.hours * f.rate
}

func totalExpenses(s []SalaryCalculator) {
	expenses := 0
	for _,v := range s {
		expenses = expenses + v.CalculateSalary()
	}
	fmt.Println("Total expenses ", expenses)
}
