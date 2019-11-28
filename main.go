package main

import (
	"github.com/andrzejd-pl/power-switch/gpio"
	"log"
	"net/http"
	"time"
)

const (
	waitTime  time.Duration = time.Second / 4
	gpioPin   int           = 17
	logFlag                 = 0
	logPrefix               = "GPIO"
)

func main() {
	http.HandleFunc("/pc/power-switch/push", pcSwitchPush)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func action(logger *log.Logger) {
	powerSwitch, err := gpio.NewGpioPowerSwitch(logger, gpioPin, waitTime)

	if err != nil {
		log.Panicln(err.Error())
	}

	defer powerSwitch.Close()

	powerSwitch.TurnOnPc()
	log.Println("Push PC switch")
}

func pcSwitchPush(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	logger := log.New(w, logPrefix, logFlag)
	action(logger)
	w.WriteHeader(http.StatusOK)
}
