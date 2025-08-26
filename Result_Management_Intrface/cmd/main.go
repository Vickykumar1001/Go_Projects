package main

import (
	"fmt"
	"log"

	"gointerfaceassessment/internal/model"
	"gointerfaceassessment/internal/service"
)

func main() {

	resultService := service.NewResultService()

	engineeringStudent := model.EngineeringStudent{
		BaseStudent: model.BaseStudent{
			ID:         "BT21EC038",
			Name:       "Vicky",
			Department: "Engineering",
		},
		Marks: map[string]int{
			"Mathematics": 85,
			"Physics":     88,
			"Programming": 92,
		},
		GradingSystem: "Credit based",
	}

	artsStudent := model.ArtsStudent{
		BaseStudent: model.BaseStudent{
			ID:         "BA21HT001",
			Name:       "Shiv",
			Department: "Arts",
		},
		Marks: map[string]int{
			"Literature": 91,
			"History":    84,
			"Philosophy": 76,
		},
		GradingSystem: "Normal",
	}

	err := resultService.AddStudent(&engineeringStudent)
	if err != nil {
		log.Printf("%v", err)
	}

	err = resultService.AddStudent(&artsStudent)
	if err != nil {
		log.Printf("%v", err)
	}

	results, err := resultService.CalculateAllResults()
	if err != nil {
		log.Printf("Error calculating results: %v", err)
	}

	for _, result := range results {
		fmt.Printf("Student: %s, Department: %s, Grade: %s, GPA: %.2f, Grading System: %s\n",
			result.Name, result.Department, result.Grade, result.GPA, result.GradingSystem)
	}
}
