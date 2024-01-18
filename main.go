package main

import (
	database "blog-service/app"
	"blog-service/controller"
	"blog-service/repository"
	"blog-service/service"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	db, err := database.NewDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo := repository.NewPostRepositoryDB(db)
	service := service.NewPostService(repo)
	controller := controller.NewPostController(service)

	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/articles/{id}", controller.GetPostByID).Methods("GET")
	r.HandleFunc("/articles", controller.GetAllPosts).Methods("GET")
	r.HandleFunc("/articles", controller.CreatePost).Methods("POST")
	r.HandleFunc("/articles/{id}", controller.UpdatePost).Methods("PUT")
	r.HandleFunc("/articles/{id}", controller.DeletePost).Methods("DELETE")

	// Use the handlers.CORS middleware to enable CORS
	// corsHandler := handlers.CORS(
	// 	handlers.AllowedOrigins([]string{"http://localhost"}),
	// 	handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"}),
	// 	handlers.AllowedHeaders([]string{"Content-Type"}),
	// )

	// Add the CORS middleware to the router
	http.Handle("/", r)

	// Listen on port 8080
	fmt.Println("Server is listening on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
