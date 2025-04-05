package main

import (
	"askme/database"
	"fmt"
	"net/http"
)

func main() {
	database.InitDB()
	database.Migrate()

	mux := http.NewServeMux()
	port := ":8080"

	mux.HandleFunc("/", MainPage)
	mux.HandleFunc("/about", AboutPage)


	http.ListenAndServe(port, mux)
}

func MainPage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write([]byte("<h1>test text</h1>"))

}
func AboutPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "<h1>About page</h1>")

}
