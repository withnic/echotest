package models

import (
	"database/sql"

	gorp "github.com/go-gorp/gorp"

	"github.com/withnic/echotest/web/app/config"
)

// initDB
func initDB() *gorp.DbMap {
	db, err := sql.Open(config.DbType, config.DbPath)
	if err != nil {
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Message{}, "Message").SetKeys(true, "id")
	dbmap.AddTableWithName(User{}, "User").SetKeys(true, "id")
	dbmap.AddTableWithName(Follow{}, "Follow").SetKeys(true, "id")

	return dbmap
}
