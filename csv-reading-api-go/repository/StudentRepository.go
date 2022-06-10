package repository

import (
	"fmt"
	"log"
	"sync"

	"code.qburst.com/dop/csv-reading-api-go/model"
	"code.qburst.com/dop/csv-reading-api-go/utils"
)


func InsertSingleStudent(student model.Student){
	db := GetDatabaseConnection()
	defer db.Close()

	insertQuery := `INSERT INTO student(id, name, college, email) VALUES ($1, $2, $3, $4)`
		_, err := db.Exec(insertQuery, student.Id, student.Name, student.College, student.Email)
		if err != nil {
			log.Fatal(err)
		}
}


func InsertMultipleStudents(wg *sync.WaitGroup, ch chan []model.Student, students ...model.Student){

	db := GetDatabaseConnection()
	defer db.Close()
	defer wg.Done() 
	

	insertQuery := `INSERT INTO student(id, name, college, email) VALUES`

	for _,student := range students {
		insertQuery = fmt.Sprintf("%v ('%v', '%v', '%v', '%v'),", insertQuery, student.Id, student.Name, student.College, student.Email)
	}

	insertQuery = insertQuery[: len(insertQuery)-1]


	_, err := db.Exec(insertQuery)
	utils.Check(err)
	ch <- students

}