package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// create a struct for student
/**
* ! struct field must uppercase for first letter for json
 */

type Student struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Address string `json:"address"`
	Job     string `json:"job"`
	Reason  string `json:"reason"`
}

// create a func based on student id from struct
func dataStudent(i int) {
	// read json file into struct
	file, err := ioutil.ReadFile("students.json")
	if err != nil {
		fmt.Println(err)
	}
	var students []Student
	json.Unmarshal(file, &students)
	// looping data from struct
	for _, student := range students {
		if student.Id == i {
			fmt.Println("Name: ", student.Name)
			fmt.Println("Address: ", student.Address)
			fmt.Println("Job: ", student.Job)
			fmt.Println("Reason: ", student.Reason)
		}
	}
}

func main() {
	// get input integer from user
	var id int
	fmt.Print("Input id: ")
	fmt.Scanln(&id)
	// call func dataStudent with id
	dataStudent(id)

}
