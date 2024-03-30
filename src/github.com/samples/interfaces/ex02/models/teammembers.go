package models

type TeamMember struct {
	FirstName  string
	FamilyName string
	Position   string
	Salary     int
}

func (f TeamMember) GetInfo() string {

	return f.FirstName + " " + f.FamilyName + ", Position: " + f.Position
}

func (f TeamMember) GetSalary() int {
	return f.Salary
}
