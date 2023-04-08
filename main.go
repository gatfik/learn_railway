package main

import (
	"dts/learn_middleware/database"
	"dts/learn_middleware/router"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	database.StartDB()
	r := router.StartServer()
	err := r.Run(":" + os.Getenv("SERVICE_PORT"))
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func ServeSample() {
	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)

	mux.Handle("/", Middleware1(endpoint))

	fmt.Println("server is running")

	err := http.ListenAndServe(":9092", mux)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func Middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("middleware 1")
		next.ServeHTTP(writer, request)
	})
}

func Middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("middleware 2")
		next.ServeHTTP(writer, request)
	})
}
