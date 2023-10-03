package config

import "time"

type DB struct {
	Host            string
	Port            int
	SslMode         string
	Name            string
	User            string
	Password        string
	Debug           bool
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifetime time.Duration
}

var db = new(DB)

func DBCfg() *DB {
	return db
}

func LoadDBCfg() {
	db.Host = GetEnvString(DB_HOST)
	db.Port = GetEnvInt(DB_PORT)
	db.User = GetEnvString(DB_USER)
	db.Password = GetEnvString(DB_PASSWORD)
	db.Name = GetEnvString(DB_NAME)
	db.SslMode = GetEnvString(DB_SSL_MODE)
	db.Debug = GetEnvBool(DB_DEBUG)
	db.MaxOpenConn = GetEnvInt(DB_MAX_OPEN_CONNECTIONS)
	db.MaxIdleConn = GetEnvInt(DB_MAX_IDLE_CONNECTIONS)

	lifetime := GetEnvInt(DB_MAX_LIFETIME_CONNECTIONS)
	db.MaxConnLifetime = time.Duration(lifetime) * time.Second
}
