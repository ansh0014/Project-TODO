package Todo

import (
	"github.com/gorilla/mux"
	"net/http"
	"your-module-name/controllers"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/todos", controllers.GetTodos).Methods(http.MethodGet)
	r.HandleFunc("/todos", controllers.AddTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", controllers.MarkDone).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods(http.MethodDelete)
	return r

}
