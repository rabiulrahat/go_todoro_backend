package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}
var tasks []Tasks


func allTasks() {

	task := Tasks{
		ID:         "1",
		TaskName:   "New projects",
		TaskDetail: "You must lead the project and finish it",
		Date:       "2024-01-22"}

	tasks = append(tasks, task)
	task1 := Tasks{
		ID:         "2",
		TaskName:   "power project",
		TaskDetail: "You must lead the project and finish it",
		Date:       "2024-01-22"}
		tasks = append(tasks, task1)
		fmt.Println("Your tasks are", tasks)

}
func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Println("I am home page")
}
func getTasks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
	fmt.Println("Heated")
}
func getTask(w http.ResponseWriter,r *http.Request){
	tasksId := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if tasksId["id"] == tasks[i].ID{
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status":"Error"})
	}
}
func createTask(w http.ResponseWriter,r *http.Request){
	fmt.Println("I am home page")
}
func deleteTask(w http.ResponseWriter,r *http.Request){
	fmt.Println("I am home page")
}
func updateTask(w http.ResponseWriter,r *http.Request){
	fmt.Println("I am home page")
}

func handleRoutes(){
	router := mux.NewRouter()
	router.HandleFunc("/",homePage).Methods("Get")
	router.HandleFunc("/gettasks",getTasks).Methods("Get")
	router.HandleFunc("/gettask/{id}",getTask).Methods("Get")
	router.HandleFunc("/createtask",createTask).Methods("POST")
	router.HandleFunc("/deletetask/{id}",deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}",updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	allTasks()
	fmt.Println("Hello")
	handleRoutes()
}
