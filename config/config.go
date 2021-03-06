package config

import (
	"database/sql"
	"strconv"
	"time"

	postgresqlpkg "github.com/agungmohmd/books-api/pkg/postgresqlPkg"
	"github.com/joho/godotenv"
)

type Configs struct {
	EnvConfig map[string]string
	DB        *sql.DB
}

func StringToInt(data string) int {
	res, err := strconv.Atoi(data)
	if err != nil {
		res = 0
	}

	return res
}

func LoadConfig() (res Configs, err error) {
	res.EnvConfig, err = godotenv.Read("../.env")
	if err != nil {
		return res, err
	}

	dbConn := postgresqlpkg.Connection{
		Host:                    res.EnvConfig["DATABASE_HOST"],
		DbName:                  res.EnvConfig["DATABASE_DB"],
		User:                    res.EnvConfig["DATABASE_USER"],
		Password:                res.EnvConfig["DATABASE_PASSWORD"],
		Port:                    StringToInt(res.EnvConfig["DATABASE_PORT"]),
		SslMode:                 res.EnvConfig["DATABASE_SSL_MODE"],
		DBMaxConnection:         StringToInt(res.EnvConfig["DATABASE_MAX_CONNECTION"]),
		DBMAxIdleConnection:     StringToInt(res.EnvConfig["DATABASE_MAX_IDLE_CONNECTION"]),
		DBMaxLifeTimeConnection: StringToInt(res.EnvConfig["DATABASE_MAX_LIFETIME_CONNECTION"]),
	}

	res.DB, err = dbConn.Connect()
	if err != nil {
		return res, err
	}
	res.DB.SetMaxIdleConns(dbConn.DBMaxConnection)
	res.DB.SetMaxIdleConns(dbConn.DBMAxIdleConnection)
	res.DB.SetConnMaxLifetime(time.Duration(dbConn.DBMaxLifeTimeConnection))

	return res, err
}
