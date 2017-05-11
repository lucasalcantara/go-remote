package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"commands"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/keyboard/space", commands.SpaceCommandEndPoint).Methods("POST")
	router.HandleFunc("/keyboard/enter", commands.EnterCommandEndPoint).Methods("POST")
	router.HandleFunc("/keyboard/left", commands.LeftCommandEndPoint).Methods("POST")
	router.HandleFunc("/keyboard/right", commands.RightCommandEndPoint).Methods("POST")
	router.HandleFunc("/keyboard/up", commands.UpCommandEndPoint).Methods("POST")
	router.HandleFunc("/keyboard/down", commands.DownCommandEndPoint).Methods("POST")
	router.HandleFunc("/keyboard/text/{input}", commands.KeyboardCommandEndPoint).Methods("POST")

	router.HandleFunc("/audio/lower", commands.LowerAudioCommandEndPoint).Methods("POST")
	router.HandleFunc("/audio/increase", commands.IncreaseAudioCommandEndPoint).Methods("POST")

	router.HandleFunc("/mouse/click", commands.ClickCommandEndPoint).Methods("POST")
	router.HandleFunc("/mouse/{x}/{y}", commands.MouseCommandEndPoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
