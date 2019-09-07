package users

import (
	"database/sql"

	"github.com/bkinnamon/passhash"
)

type User struct {
	ID    int64
	Email string
	hash  string
	Name  string
}

func createUser(db *sql.DB, email string, password string, name string) *User {
	sql := "INSERT INTO users (email, hash, name) VALUES ($1, $2, $3)"
	stmt, _ := db.Prepare(sql)

	hash, _ := passhash.GenerateFromPassword(password)

	result, _ := stmt.Exec(email, hash, name)

	id, _ := result.LastInsertId()

	sql = "SELECT id, email, name FROM users WHERE id = $1"
	stmt, _ = db.Prepare(sql)

	var u *User
	_ = stmt.QueryRow(id).Scan(&u.ID, &u.Email, &u.Name)

	return u
}
