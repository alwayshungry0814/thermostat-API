package main

import (
	"encoding/json"
	"net/http"
)

// GET request for the current temperature
func getTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(temperatures)
}

// GET request for indicating cooling or heating
func getthermostatState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thermostatStates)
}

// GET request for current temperature settings for both heating and cooling
func getTemperatureSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(temperatureSettings)
}

// PUT request for changing the temperature setting
func changeTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for index := range temperatureSettings {
		temperatureSettings = append(temperatureSettings[:index], temperatureSettings[index+1:]...)
		var temperatureSetting TemperatureSetting
		_ = json.NewDecoder(r.Body).Decode(&temperatureSetting)
		temperatureSettings = append(temperatureSettings, temperatureSetting)
		json.NewEncoder(w).Encode(temperatureSetting)
		return
	}
}

// TODO: for the Bonus stretch goal (actual temp respond slowly to state of the thermostat)
// after putting in a request to do the PUT, i would change the thermostatState accordingly
// then add a timer to increment temperature based on a timer until it reached the desired temp

// TODO: add custom error handling for PUT request body having proper key values.
// error should indicate what key values should be on body

// TODO: add unit tests for getTemperature, getthermostatState, getTemperatureSetting, changeTemperature
// test that returns proper key values
