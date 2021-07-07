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

	employees := []SalaryCalculator{p1, p2, c1}
	totalExpenses(employees)
}

func (p permanent) CalculateSalary() int {
	return p.basic + p.bonus
}

func (c contract) CalculateSalary() int {
	return c.basic
}

func totalExpenses(s []SalaryCalculator) {
	expenses := 0
	for _,v := range s {
		expenses = expenses + v.CalculateSalary()
	}
	fmt.Println("Total expenses ", expenses)
}
