package main

import (
	"ex01/interfaces"
	"ex01/models"
	"fmt"
)

func main() {
	fmt.Println("start sample EX01")
	lead01 := models.TeamLead{
		FirstName:  "Lead01",
		FamilyName: "TeamLead",
		TeamSize:   30,
		Salary:     200,
	}
	member01 := models.TeamMember{
		FirstName:  "Member01",
		FamilyName: "Team",
		Position:   "developer",
		Salary:     2000,
	}

	var teamsInterface interfaces.Employee
	teamsInterface = lead01
	a := teamsInterface.GetInfo()
	b := teamsInterface.GetSalary()
	models.PrintInfo(a, b)
	teamsInterface = member01
	a = teamsInterface.GetInfo()
	b = teamsInterface.GetSalary()
	models.PrintInfo(a, b)

}
