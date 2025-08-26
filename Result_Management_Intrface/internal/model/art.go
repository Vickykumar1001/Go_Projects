package model

type ArtsStudent struct {
	BaseStudent
	Marks         map[string]int
	GradingSystem string
}

func (s *ArtsStudent) GetID() string {
	return s.ID
}

func (s *ArtsStudent) GetName() string {
	return s.Name
}

func (s *ArtsStudent) GetDepartment() string {
	return s.Department
}
func (s *ArtsStudent) GetGradingSystem() string {
	return s.GradingSystem
}

func (s *ArtsStudent) CalculateGrade() (string, float64, error) {
	totalMarks := 0

	for _, mark := range s.Marks {
		totalMarks += mark
	}

	average := float64(totalMarks) / float64(len(s.Marks))

	gpa := average / 100.0 * 10.0

	var grade string
	switch {
	case average >= 85:
		grade = "A+"
	case average >= 75:
		grade = "A"
	case average >= 65:
		grade = "B"
	case average >= 55:
		grade = "C"
	case average >= 45:
		grade = "D"
	default:
		grade = "F"
	}

	return grade, gpa, nil
}
