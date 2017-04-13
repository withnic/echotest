package models

import (
	"database/sql"
	"log"

	gorp "github.com/go-gorp/gorp"

	"github.com/withnic/echotest/web/app/config"
)

// MessageWithUser is include relationships
type MessageWithUser struct {
	Mes  Message
	User User
}

func initMessageWithUserDB() *gorp.DbMap {
	db, err := sql.Open(config.DbType, config.DbPath)
	if err != nil {
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Message{}, "Message").SetKeys(true, "id")
	dbmap.AddTableWithName(User{}, "User").SetKeys(true, "id")
	return dbmap
}

// GetAll is get MessageWithUser
func (mes *MessageWithUser) GetAll() []MessageWithUser {
	var messages []MessageWithUser
	dbmap := initMessageWithUserDB()
	rows, err := dbmap.Query("select * from Message inner join User on User.id = Message.user_id order by Message.created_at desc")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var mes Message
		var user User
		rows.Scan(
			&mes.ID,
			&mes.Body,
			&mes.UserID,
			&mes.CreatedAt,
			&user.ID,
			&user.Email,
			&user.Passwd,
		)
		messages = append(messages, MessageWithUser{
			Mes:  mes,
			User: user,
		})
	}

	return messages
}
