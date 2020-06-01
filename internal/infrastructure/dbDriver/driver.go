package dbDriver

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stetsd/micro-brick/config"
)

type DbDriver struct {
	dialect string
	Db      *sqlx.DB
}

var dbDriver DbDriver

func NewDbDriver(conf config.Config) (*DbDriver, error) {

	if dbDriver.dialect != "" {
		return &dbDriver, nil
	}

	connString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Get(config.DbHost),
		conf.Get(config.DbPort),
		conf.Get(config.DbUser),
		conf.Get(config.DbPass),
		conf.Get(config.DbName),
	)

	db, err := sqlx.Connect("postgres", connString)

	if err != nil {
		return nil, err
	}

	dbDriver = DbDriver{
		dialect: "postgres",
		Db:      db,
	}

	return &dbDriver, nil
}
