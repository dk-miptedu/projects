package models

import "strconv"

type TeamLead struct {
	FirstName  string
	FamilyName string
	TeamSize   int
	Salary     int
}

func (f TeamLead) GetInfo() string {
	return f.FirstName + " " + f.FamilyName + ", TeamSize: " + strconv.Itoa(f.TeamSize)
}

func (f TeamLead) GetSalary() int {
	return f.Salary
}
