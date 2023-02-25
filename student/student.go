package student

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Msv         string `json:"msv"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Age         int    `json:"age"`
}

func readStudentsFromFile() ([]Student, error) {
	file, err := os.Open("students.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var students []Student
	err = json.NewDecoder(file).Decode(&students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func writeStudentsToFile(students []Student) error {
	file, err := os.Create("students.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(students)
	if err != nil {
		return err
	}

	return nil
}

func CreateStudent(s Student) error {
	students, err := readStudentsFromFile()
	if err != nil {
		return err
	}

	for _, st := range students {
		if st.Msv == s.Msv {
			return fmt.Errorf("student with msv %s already exists", s.Msv)
		}
	}

	students = append(students, s)
	err = writeStudentsToFile(students)
	if err != nil {
		return err
	}

	return nil
}

func GetStudent(msv string) (*Student, error) {
	students, err := readStudentsFromFile()
	if err != nil {
		return nil, err
	}

	for _, s := range students {
		if s.Msv == msv {
			return &s, nil
		}
	}

	return nil, fmt.Errorf("student with msv %s not found", msv)
}

func GetAllStudents() ([]Student, error) {
	return readStudentsFromFile()
}

func UpdateStudent(s Student) error {
	students, err := readStudentsFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, st := range students {
		if st.Msv == s.Msv {
			students[i] = s
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("student with msv %s not found", s.Msv)
	}

	err = writeStudentsToFile(students)
	if err != nil {
		return err
	}

	return nil
}

func DeleteStudent(msv string) error {
	students, err := readStudentsFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, s := range students {
		if s.Msv == msv {
			students = append(students[:i], students[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("student with msv %s not found", msv)
	}

	err = writeStudentsToFile(students)
	if err != nil {
		return err
	}

	return nil
}
