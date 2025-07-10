package routes

import (
	"github.com/gorilla/mux"
	"net/http"
  "Todo/controllers"
  "Todo/middleware"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	
	r.HandleFunc("/Register", controllers.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	// Secure group
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)

	r.HandleFunc("/todos", controllers.GetTodo).Methods(http.MethodGet)
	r.HandleFunc("/todos", controllers.AddTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", controllers.MarkDone).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods(http.MethodDelete)
	return r

}
