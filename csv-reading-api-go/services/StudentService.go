package services

import (
	"log"
	"mime/multipart"
	"sync"

	"code.qburst.com/dop/csv-reading-api-go/model"
	"code.qburst.com/dop/csv-reading-api-go/repository"
	"code.qburst.com/dop/csv-reading-api-go/utils"
)

var wg sync.WaitGroup

func CreateStudentsUsingCsvWithoutGoRoutines(file multipart.File) ([]model.Student, error) {

	data, err := utils.ReadCsvData(file)

	if err != nil {
		log.Fatal(err)
	}

	var students []model.Student

	for _, line := range data {

		students = append(students, model.Student{
			Id:      line[0],
			Name:    line[1],
			College: line[2],
			Email:   line[3],
		})
	}

	db := repository.GetDatabaseConnection()
	defer db.Close()

	for _, student := range students {
		repository.InsertSingleStudent(student)
	}

	return students, nil
}

func CreateStudentsUsingCsvAndGoRoutines(file multipart.File) ([]model.Student, error) {

	var students []model.Student
	const batchSize int = 100

	ch := make(chan []model.Student, 1002)


	data, err := utils.ReadCsvData(file)
	utils.Check(err)

	for _, line := range data {
		students = append(students, model.Student{
			Id:      line[0],
			Name:    line[1],
			College: line[2],
			Email:   line[3],
		})
	}

	// Logic for sending batchwise request to postgres

	for i := 0; i < len(students); i += batchSize {
		end := i + batchSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(students) {
			end = len(students)
		}

		wg.Add(1)
		go repository.InsertMultipleStudents(&wg, ch, students[i:end]...)
	}
	
	wg.Wait()
	close(ch)
	var result []model.Student
	for v := range ch{
		result = append(result, v...)
	}
	return result, nil
}

func GetAllStudents() ([]model.Student, error) {
	selectString := `SELECT * FROM student`

	db := repository.GetDatabaseConnection()
	defer db.Close()

	rows, err := db.Query(selectString)

	if err != nil {
		log.Fatal(err)
	}

	var students []model.Student

	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.Id, &student.Name, &student.College, &student.Email)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}
	return students, nil
}
