package model

type Info struct {
	NameCPU        string   `json:"NameCPU,omitempty"`
	TemperatureCPU []string `json:"TemperatureCPU,omitempty"`
	NameGPU        string   `json:"NameGPU,omitempty"`
	TemperatureGPU string   `json:"TemperatureGPU,omitempty"`
	NameHDD        string   `json:"NameHDD,omitempty"`
	TemperatureHDD string   `json:"TemperatureHDD,omitempty"`
	LocalIPAddress string   `json:"LocalIpAddress,omitempty"`
	MACaddress     string   `json:"MACaddress,omitempty"`
}
