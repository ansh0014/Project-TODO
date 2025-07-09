package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	
	"Todo/model"
	

	
)

var Todos []model.Todo
var NextID=1
func GetTodo(w http.ResponseWriter, r*http.Request)  {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Todos)
	
}
func AddTodo(w http.ResponseWriter, r*http.Request){
   var todo model.Todo
   json.NewDecoder(r.Body).Decode(&todo)
   todo.ID=NextID
   NextID++
   Todos=append(Todos, todo)
   json.NewEncoder(w).Encode(todo)
}
func MarkDone(w http.ResponseWriter, r* http.Request){
id,_:=strconv.Atoi(mux.Vars(r)["id"])
	for i := range Todos {
		if Todos[i].ID == id {
			Todos[i].Done = true
			json.NewEncoder(w).Encode(Todos[i])
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i := range Todos {
		if Todos[i].ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			w.Write([]byte("Deleted"))
			return
		}
	}
	http.NotFound(w, r)
}



