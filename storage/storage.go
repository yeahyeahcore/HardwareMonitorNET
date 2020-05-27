package storage

import (
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/HardwareMonitorNET/config"
)

var (
	db  *sqlx.DB
	log *logrus.Logger

	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

// Tables
var (
	Devices devices
)

func Init() {
	fmt.Println("\nInit storage...")

	connCfg, err := pgx.ParseURI(config.Server.Storage.Connection)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connCfg,
		MaxConnections: 20,
		AcquireTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	native := stdlib.OpenDBFromPool(pool)
	db = sqlx.NewDb(native, config.Server.Storage.Driver)

	db.SetMaxIdleConns(config.Server.Storage.MaxIdleConnection)
	db.SetMaxOpenConns(config.Server.Storage.MaxOpenConnection)

	fmt.Println("Done.")
}

func Tx() (*sqlx.Tx, error) {
	return db.Beginx()
}

func exec(tx *sqlx.Tx, stmt squirrel.Sqlizer) error {
	query, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	if tx == nil {
		_, err = db.Exec(query, args...)
	} else {
		_, err = tx.Exec(query, args...)
	}
	return err
}
