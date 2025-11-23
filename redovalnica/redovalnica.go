// Package redovalnica provides utilities for managing students and their grades.
//
// The package allows you to initialize students, assign grades, calculate averages,
// and print grade reports. It is designed as a simple utility for working with
// redovalnica (grade books) using Go maps to store students by their vpisna številka.
//
// Example usage:
//
//	studenti := make(map[string]redovalnica.Student)
//
//	var s redovalnica.Student
//	redovalnica.InitStudent(&s, "Ana", "Novak")
//	studenti["63230001"] = s
//
//	redovalnica.DodajOceno(studenti, "63230001", 9)
//	redovalnica.DodajOceno(studenti, "63230001", 10)
//
//	redovalnica.IzpisRedovalnice(studenti)
//	redovalnica.IzpisiKoncniUspeh(studenti)
//
// The grade book supports the following features:
//   - Initializing new students with empty grade lists
//   - Adding grades in range 0–10
//   - Computing averages with at least six grades
//   - Printing detailed and final grade summaries
package redovalnica

import (
	"fmt"
)

// Student represents a student with a first name, last name, and a list of grades.
type Student struct {
	ime     string
	priimek string
	ocene   []int
}

// DodajOceno adds a grade to a student identified by the vpisna številka.
// Grades outside the range 0–10 are ignored.
// If the student does not exist, an error message is printed.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena > 10 || ocena < 0 {
		return
	}
	if s, ok := studenti[vpisnaStevilka]; !ok {
		fmt.Println("Student ni na seznamu")
		return
	} else {
		s.ocene = append(s.ocene, ocena)
		studenti[vpisnaStevilka] = s
	}
}

// povprecje calculates a student's average grade.
// If the student does not exist, it returns -1.
// If the student has fewer than 6 grades, it returns 0.
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	if s, ok := studenti[vpisnaStevilka]; !ok {
		fmt.Println("Student ni na seznamu")
		return -1.0
	} else {
		var sum = 0
		var l = len(s.ocene)
		if l < 6 {
			return 0.0
		}
		for _, v := range s.ocene {
			sum += v
		}
		return float64(sum) / float64(l)
	}
}

// IzpisRedovalnice prints all students with their grades.
func IzpisRedovalnice(studenti map[string]Student) {
	for k, s := range studenti {
		fmt.Printf("%s - %s %s: %v\n", k, s.ime, s.priimek, s.ocene)
	}
}

// IzpisiKoncniUspeh prints each student's average grade and a verbal assessment.
func IzpisiKoncniUspeh(studenti map[string]Student) {
	for k, s := range studenti {
		povp := povprecje(studenti, k)
		fmt.Printf("%s %s: povprečna ocena %g -> %s\n",
			s.ime, s.priimek, povp, getUspeh(povp))
	}
}

// getUspeh returns a text description of a student's performance based on the average grade.
func getUspeh(povpOcena float64) string {
	if povpOcena >= 9 {
		return "Odličen študent!"
	} else if povpOcena >= 6 {
		return "Povprečen študent"
	}
	return "Neuspešen študent"
}

// InitStudent initializes a new Student with the provided first and last name.
// The student starts with an empty grade list.
func InitStudent(s *Student, ime string, priimek string) {
	s.ime = ime
	s.priimek = priimek
	s.ocene = []int{}
}
