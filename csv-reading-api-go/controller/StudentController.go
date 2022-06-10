package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"code.qburst.com/dop/csv-reading-api-go/model"
	"code.qburst.com/dop/csv-reading-api-go/services"
	"github.com/gorilla/mux"
)



func HandleRequests(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/csv/api/students", getAllStudents).Methods("GET")
	router.HandleFunc("/csv/api/studentswithgoroutines", createStudentsUsingCsvAndGoRoutines).Methods("POST")
	router.HandleFunc("/csv/api/student", createStudentsUsingCsvWIthoutGoRoutines).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))	
}


// creating students using csv file

func createStudentsUsingCsvWIthoutGoRoutines(w http.ResponseWriter, r *http.Request){
	file, _, err := r.FormFile("filename")
	if err != nil {
		log.Println("error: createStudentsUsingCsvWIthoutGoRoutines:29")
		log.Fatal(err)
	}
	students, err := services.CreateStudentsUsingCsvWithoutGoRoutines(file)

	if err != nil {
		log.Println("error: createStudentsUsingCsvWIthoutGoRoutines:36")
		log.Fatal(err)
	}

	response := model.JsonResponse{
		Type: "Success",
		Data: students,
		Message: "Data inserted successfully",
	}
	json.NewEncoder(w).Encode(response)
	
}

// create students in bulk using csv file

func createStudentsUsingCsvAndGoRoutines(w http.ResponseWriter, r *http.Request){
	file, _, err := r.FormFile("filename")
	if err != nil {
		fmt.Println("error: createStudentsUsingCsvAndGoRoutines:54")
		log.Fatal(err)
	}
	students, err := services.CreateStudentsUsingCsvAndGoRoutines(file)

	if err != nil {
		fmt.Println("error: createStudentsUsingCsvAndGoRoutines:60")
		log.Fatal(err)
	}

	response := model.JsonResponse{
		Type: "Success",
		Data: students,
		Message: "Data inserted successfully",
	}
	json.NewEncoder(w).Encode(response)
	
}


// getting all students

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := services.GetAllStudents()
	if err != nil {
		log.Fatal(err)
	}
	response := model.JsonResponse{
		Type: "Success",
		Data: students,
		Message: "Data retrieved successfully",
	}
	json.NewEncoder(w).Encode(response)	
}


