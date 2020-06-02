package storage

import "github.com/jmoiron/sqlx"

//Device структура из БД(и вот тут)
type Device struct {
	ID         string  `db:"id" json:"id,omitempty"`
	PC         string  `db:"pc_name" json:"pc,omitempty"`
	MACAddress string  `db:"mac_address" json:"mac_address,omitempty"`
	CPU        *string `db:"cpu_name" json:"cpu_name,omitempty"`
	HDD        *string `db:"hdd_name" json:"hdd_name,omitempty"`
	GPU        *string `db:"gpu_name" json:"gpu_name,omitempty"`
}

type devices struct{}

func (d *devices) Insert(tx *sqlx.Tx, device *Device) error {
	_, err := tx.Exec(`
		INSERT INTO devices(
			id,
			pc_name,
			mac_address,
			cpu_name,
			hdd_name,
			gpu_name
		)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		device.ID,
		device.PC,
		device.MACAddress,
		device.CPU,
		device.HDD,
		device.GPU,
	)
	return err
}
