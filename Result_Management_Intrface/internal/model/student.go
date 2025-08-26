package model

type Student interface {
	GetID() string
	GetName() string
	GetDepartment() string
	GetGradingSystem() string
	CalculateGrade() (string, float64, error)
}

type BaseStudent struct {
	ID         string
	Name       string
	Department string
}

func (s BaseStudent) GetID() string {
	return s.ID
}

func (s BaseStudent) GetName() string {
	return s.Name
}

func (s BaseStudent) GetDepartment() string {
	return s.Department
}

type Result struct {
	BaseStudent
	Grade         string
	GPA           float64
	GradingSystem string
}
