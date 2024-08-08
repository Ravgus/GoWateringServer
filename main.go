package main

import (
	"GoWateringServer/internal"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(404), http.StatusNotFound)
}

func getInstructionsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "watering-instruction.json")
}

func getHistoryHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "history.json")
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	internal.AddJsonData("history.json", internal.GenerateDate())

	http.Redirect(w, r, "/", 301)
}

func main() {
	internal.LoadEnv()

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/getInstructions", getInstructionsHandler).Methods("GET")
	router.HandleFunc("/getHistory", getHistoryHandler).Methods("GET")
	router.HandleFunc("/notify", notifyHandler).Methods("POST")

	http.Handle("/", router)

	port := os.Getenv("SERVER_PORT")
	if len(port) == 0 {
		log.Fatal("SERVER_PORT is empty!")
	}

	fmt.Println("Ready to go!")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
