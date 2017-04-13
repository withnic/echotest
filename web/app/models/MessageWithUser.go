package models

import "log"

// MessageWithUser is include relationships
type MessageWithUser struct {
	Mes  Message
	User User
}

// GetAll is get MessageWithUser
func (mes *MessageWithUser) GetAll() []MessageWithUser {
	var messages []MessageWithUser
	dbmap := initDB()
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
