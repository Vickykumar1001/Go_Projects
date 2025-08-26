package service

import (
	"errors"
	"fmt"
	"gointerfaceassessment/internal/model"
)

var (
	ErrNoStudents      = "no students registered"
	ErrDuplicateID     = "student with this ID already exists"
	ErrNoStudentWithID = "student not found with ID"
	ErrCalcGrade       = "failed to calculate grade for student"
)

type ResultService struct {
	students []model.Student
}

func NewResultService() *ResultService {
	return &ResultService{
		students: make([]model.Student, 0),
	}
}

func (s *ResultService) AddStudent(student model.Student) error {
	for _, existStudent := range s.students {
		if existStudent.GetID() == student.GetID() {
			return fmt.Errorf("%w: %s", errors.New(ErrDuplicateID), student.GetID())
		}
	}
	s.students = append(s.students, student)
	return nil
}

func (s *ResultService) GetStudentByID(id string) (model.Student, error) {
	for _, student := range s.students {
		if student.GetID() == id {
			return student, nil
		}
	}
	return nil, fmt.Errorf("%w: %s", errors.New(ErrNoStudentWithID), id)
}

func (s *ResultService) CalculateAllResults() ([]model.Result, error) {
	if len(s.students) == 0 {
		return nil, errors.New(ErrNoStudents)
	}

	results := make([]model.Result, 0, len(s.students))

	for _, student := range s.students {
		grade, gpa, err := student.CalculateGrade()
		if err != nil {
			return nil, fmt.Errorf("%w %s: %w", errors.New(ErrCalcGrade), student.GetName(), err)
		}

		result := model.Result{
			BaseStudent: model.BaseStudent{
				ID:         student.GetID(),
				Name:       student.GetName(),
				Department: student.GetDepartment(),
			},
			Grade:         grade,
			GPA:           gpa,
			GradingSystem: student.GetGradingSystem(),
		}

		results = append(results, result)
	}

	return results, nil
}
