package main

import (
	"fmt"
)

type SalaryCalculator interface {
	calculateSalary() int
}

type Permanent struct {
	empId int
	basic int
	pf int
}

type Contract struct {
	empId int
	basic int
}

func (p Permanent) calculateSalary() int {
	return p.basic + p.pf
}

func (c Contract) calculateSalary() int {
	return c.basic
}

func totalExpenses(s []SalaryCalculator) {
	expense := 0
	for _, v:= range s {
		fmt.Printf("emp:%T\n", v)
		expense = expense + v.calculateSalary()
	}
	fmt.Printf("Total exp:%d", expense)
}

func main(){
	pEmp1 := Permanent{101,1000,20}
	pEmp2 := Permanent{102,2000,30}
	cEmp1 := Contract{202,3000}
	
	emps := []SalaryCalculator{pEmp1,pEmp2,cEmp1}
	totalExpenses(emps)	
}

