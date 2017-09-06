package models

import (
	"github.com/wartent/test/db"
)

type User struct {
	ID        int    `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

type UserModel struct{}

func (m UserModel) One(userID int64) (user User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM public.user WHERE id=$1", userID)

	if err != nil {
		return user, err
	}

	return user, nil
}
