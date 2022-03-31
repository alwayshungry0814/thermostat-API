package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// temporary variables to hold seed data. TODO: create a db
var temperatures []Temperature
var temperatureSettings []TemperatureSetting
var thermostatStates []ThermostatState

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// seed data. TODO: add to database
	temperatures = append(temperatures, Temperature{Temperature: 67})
	temperatureSettings = append(temperatureSettings, TemperatureSetting{CoolTemp: 72, HeatTemp: 68})
	thermostatStates = append(thermostatStates, ThermostatState{Cooling: false, Heating: true})

	// Route handles & endpoints
	r.HandleFunc("/temperature", getTemperature).Methods("GET")
	r.HandleFunc("/temperatureSetting", getTemperatureSetting).Methods("GET")
	r.HandleFunc("/thermostatState", getthermostatState).Methods("GET")
	r.HandleFunc("/temperatureSetting", changeTemperature).Methods("PUT")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
