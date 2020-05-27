package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

//CPU структура из БД
type CPU struct {
	Temp  []int     `db:"cpu_temp" json:"temp,omitempty"`
	Clock []float32 `db:"cpu_clock" json:"clock,omitempty"`
	Load  []int     `db:"cpu_load" json:"load,omitempty"`
}

//Power структура из БД
type Power struct {
	CPUPackage  *float32 `db:"power_cpu_package" json:"cpu_package,omitempty"`
	CPUCores    *float32 `db:"power_cpu_cores" json:"cpu_cores,omitempty"`
	CPUGraphics *float32 `db:"power_cpu_graphics" json:"cpu_graphics,omitempty"`
	CPUDram     *float32 `db:"power_cpu_dram" json:"cpu_dram,omitempty"`
}

//Memory структура из БД
type Memory struct {
	Load      *float32 `db:"memory_load" json:"load,omitempty"`
	Used      *float32 `db:"memory_used" json:"used,omitempty"`
	Available *float32 `db:"memory_available" json:"available,omitempty"`
}

//HDD структура из БД
type HDD struct {
	Temp *int `db:"hdd_temp" json:"temp,omitempty"`
}

//GPU структура из БД
type GPU struct {
	Clocks     *float32 `db:"gpu_clocks" json:"clocks,omitempty"`
	Load       *float32 `db:"gpu_load" json:"load,omitempty"`
	MemoryUsed *float32 `db:"gpu_memory_used" json:"memory_used,omitempty"`
	MemoryFree *float32 `db:"gpu_memory_free" json:"memory_free,omitempty"`
}

//Parameter ето крч структура данных из БД
type Parameter struct {
	ID        uint64    `db:"id" json:"id,omitempty"`
	DeviceID  string    `db:"device_id" json:"device_id,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`

	CPU    `json:"cpu,omitempty"`
	Power  `json:"power,omitempty"`
	Memory `json:"memory,omitempty"`
	HDD    `json:"hdd,omitempty"`
	GPU    `json:"gpu,omitempty"`
}

type parameters struct{}

func (p *parameters) Insert(tx *sqlx.Tx, param *Parameter) error {
	_, err := tx.Exec(`
		INSERT INTO parameters(
			id, 
			device_id, 
			cpu_temp,
			cpu_clock,
			cpu_load,
			power_cpu_package,
			power_cpu_cores,
			power_cpu_graphics,
			power_cpu_dram,
			memory_load,
			memory_used,
			memory_available,
			hdd_temp,
			gpu_clocks,
			gpu_load, 
			gpu_memory_used,
			gpu_memory_free
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`,
		param.ID,
		param.DeviceID,
		param.CPU.Temp,
		param.CPU.Clock,
		param.CPU.Load,
		param.Power.CPUPackage,
		param.Power.CPUCores,
		param.Power.CPUGraphics,
		param.Power.CPUDram,
		param.Memory.Load,
		param.Memory.Used,
		param.Memory.Available,
		param.HDD.Temp,
		param.GPU.Clocks,
		param.GPU.Load,
		param.GPU.MemoryUsed,
		param.GPU.MemoryFree,
	)
	return err
}
