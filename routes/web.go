package routes

import (
	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/controllers"
)

func Init() *mux.Router {
	// Router ini menggunakan gorilla mux
	route := mux.NewRouter()

	// Variable controllers ini mengambil dari file controllers/todo.go
	route.HandleFunc("/", controllers.Show)
	route.HandleFunc("/update/{id}", controllers.Update)
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete)
	route.HandleFunc("/complete/{id}", controllers.Complete)
	route.HandleFunc("/uncomplete/{id}", controllers.UnComplete)

	/* TODO: membuat fungsi untuk mengupdate data [Menerima dari url /update/{id}]  */

	return route
}
