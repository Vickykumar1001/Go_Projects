package model

type EngineeringStudent struct {
	BaseStudent
	Marks         map[string]int
	GradingSystem string
}

func (s *EngineeringStudent) GetID() string {
	return s.ID
}

func (s *EngineeringStudent) GetName() string {
	return s.Name
}

func (s *EngineeringStudent) GetDepartment() string {
	return s.Department
}
func (s *EngineeringStudent) GetGradingSystem() string {
	return s.GradingSystem
}

func (s *EngineeringStudent) CalculateGrade() (string, float64, error) {
	totalMarks := 0
	totalWeight := 0

	weights := map[string]int{
		"Mathematics": 3,
		"Physics":     2,
		"Programming": 3,
	}

	for subject, mark := range s.Marks {
		if weight, exists := weights[subject]; exists {
			totalMarks += mark * weight
			totalWeight += weight
		}
	}

	average := float64(totalMarks) / float64(totalWeight)

	gpa := average / 100.0 * 10.0

	var grade string
	switch {
	case average >= 90:
		grade = "A+"
	case average >= 80:
		grade = "A"
	case average >= 70:
		grade = "B"
	case average >= 60:
		grade = "C"
	case average >= 50:
		grade = "D"
	default:
		grade = "F"
	}

	return grade, gpa, nil
}
