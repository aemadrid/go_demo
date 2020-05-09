package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", JsonHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	str := `{` +
		`"version":3,` +
		`"status":"ok",` +
		`"message":"Hello World",` +
		`"language":"go",` +
		`"env":{` +
		`"pod_name":"` + os.Getenv("POD_NAME") + `"` +
		`}` +
		`}`
	io.WriteString(w, str)
}
