package storage

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/lib/pq"
)

//CPU структура из БД
type CPU struct {
	Temp  []float32 `db:"cpu_temp" json:"temp,omitempty"`
	Clock []float32 `db:"cpu_clock" json:"clock,omitempty"`
}

//Memory структура из БД(вот тут поправить нужно будет)
type Memory struct {
	Load      *float32 `db:"memory_load" json:"load,omitempty"`
	Used      *float32 `db:"memory_used" json:"used,omitempty"`
	Available *float32 `db:"memory_available" json:"available,omitempty"`
}

//HDD структура из БД
type HDD struct {
	Temp *float32 `db:"hdd_temp" json:"temp,omitempty"`
}

//GPU структура из БД
type GPU struct {
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
			created_at,
			cpu_temp,
			cpu_clock,
			memory_load,
			memory_used,
			memory_available,
			hdd_temp,
			gpu_load, 
			gpu_memory_used,
			gpu_memory_free
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		param.ID,
		param.DeviceID,
		param.CreatedAt,
		pq.Array(param.CPU.Temp),
		pq.Array(param.CPU.Clock),
		param.Memory.Load,
		param.Memory.Used,
		param.Memory.Available,
		param.HDD.Temp,
		param.GPU.Load,
		param.GPU.MemoryUsed,
		param.GPU.MemoryFree,
	)
	return err
}
