package config

//omitempy если ноль будет то ничего не запишет
type config struct {
	Server server `json:"server,omitempty"`
	Client client `json:"client,omitempty"`
}

type server struct {
	Host    string   `json:"host,omitempty"`
	Port    string   `json:"port,omitempty"`
	Storage *storage `json:"storage,omitempty"`
}

type client struct {
	ID       string `json:"id,omitempty"`
	PauseSec int    `json:"pause_sec,omitempty"`
}

type storage struct {
	Driver            string `json:"driver,omitempty"`
	Connection        string `json:"connection,omitempty"`
	MaxIdleConnection int    `json:"max_idle_connection,omitempty"`
	MaxOpenConnection int    `json:"max_open_connection,omitempty"`
}
