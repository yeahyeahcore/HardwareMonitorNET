package config

type config struct {
	Server server `json:"server,omitempty"`
	Client client `json:"client,omitempty"`
}

type server struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}

type client struct {
	Time string `json:"time,omitempty"`
}
