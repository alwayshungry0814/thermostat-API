package main

type Temperature struct {
	Temperature int `json:"temperature"`
}

type TemperatureSetting struct {
	CoolTemp int `json:"coolTemp"`
	HeatTemp int `json:"heatTemp"`
}

type ThermostatState struct {
	Cooling bool `json:"cooling"`
	Heating bool `json:"heating"`
}

// these structs were based on requirements that the API user would need
// the struct could have been combined into one but for the purpose of meeting requirements
// in sequential steps I used the structs above.

// I could have also used the struct below.

// type ThermostatSetting struct {
// 	CurrentTemperature int                 `json:"currentTemperature`
// 	TemperatureSetting *TemperatureSetting `json:"temperatureSetting"`
// 	ThermostatState    *ThermostatState    `json:"thermostatState"`
// }

// type TemperatureSetting struct {
// 	CoolTemp int `json:"coolTemp"`
// 	HeatTemp int `json:"heatTemp"`
// }

// type ThermostatState struct {
// 	Cooling bool `json:"cooling"`
// 	Heating bool `json:"heating"`
// }
