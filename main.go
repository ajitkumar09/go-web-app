package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func newRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", homePage)
	mux.HandleFunc("GET /home", homePage)
	mux.HandleFunc("GET /courses", coursePage)
	mux.HandleFunc("GET /about", aboutPage)
	mux.HandleFunc("GET /contact", contactPage)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	return mux
}

func main() {
	err := http.ListenAndServe("0.0.0.0:8080", newRouter())
	if err != nil {
		log.Fatal(err)
	}
}
