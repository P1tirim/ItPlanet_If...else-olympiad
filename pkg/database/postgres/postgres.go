package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var errPgConfigError = errors.New("empty db config pointer")

// Config type.
type Config struct {
	Host         string `yaml:"host" json:"host" envconfig:"API_DB_HOST"`
	Port         int    `yaml:"port" json:"port" envconfig:"API_DB_PORT"`
	Dbname       string `yaml:"dbname" json:"dbname" envconfig:"API_DB_NAME"`
	User         string `yaml:"user" json:"user" envconfig:"API_DB_USER"`
	Password     string `yaml:"password" json:"password" envconfig:"API_DB_PASSWORD"`
	Sslmode      string `yaml:"sslmode" json:"sslmode" envconfig:"API_DB_SSLMODE"`
	Poolmaxconns int    `yaml:"poolmaxconns" json:"poolmaxconns" envconfig:"API_DB_POOLMAXXCONN"`
}

// MakePostgres make postgres repo.
func MakePostgres(conf *Config) (*pgxpool.Pool, error) {
	if conf == nil {
		return nil, errPgConfigError
	}

	DBUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&pool_max_conns=%d&pool_max_conn_lifetime=1s&pool_max_conn_idle_time=1s&pool_health_check_period=1s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Dbname,
		conf.Sslmode,
		conf.Poolmaxconns)

	confPgx, err := pgxpool.ParseConfig(DBUrl)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(context.Background(), confPgx)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}
