package commands

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/go-vgo/robotgo"
)

const speed = 8

func MouseCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	moveX, err := strconv.Atoi(params["x"])
	if err != nil {
		return
	}

	moveY, err := strconv.Atoi(params["y"])
	if err != nil {
		return
	}

	x, y := robotgo.GetMousePos()
	robotgo.MoveMouse(x+moveX*speed, y+moveY*speed)
}

func ClickCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	robotgo.MouseClick()
}

func KeyboardCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, input := range params {
		for _, letter := range input {
			robotgo.KeyTap(string(letter))
		}
	}
}

func SpaceCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	robotgo.KeyTap("escape")
}

func EnterCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	robotgo.KeyTap("enter")
}

func LeftCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	robotgo.KeyTap("left")
}

func RightCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	robotgo.KeyTap("right")
}

func LowerAudioCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	robotgo.KeyTap("audio_vol_down")
}

func UpCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	robotgo.KeyTap("up")
}

func DownCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	robotgo.KeyTap("down")
}

func IncreaseAudioCommandEndPoint(w http.ResponseWriter, req *http.Request) {
	robotgo.KeyTap("audio_vol_up")
}
