package internal

import "time"

type MysqlConf struct {
	Host         string `toml:"Host"`
	Port         int
	User         string
	Password     string
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
}

type DBConf struct {
	App     string                    `toml:"APP"`
	EnvConf map[string]MysqlConf `toml:"ENV"`
}

func (d DBConf)Set(s string)error{
	return nil
}

type TableConf struct {
	Name string
}
