package mysql

import (
	"database/sql"
	"fmt"
	entity "suggestApp/enity"
)

func (d *MySqlDB) IsUniquePhoneNumber(phoneNumber string) (bool, error) {
	var user entity.User
	var created_at []uint8

	row := d.db.QueryRow("select * from users where phone_number= ?", phoneNumber)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &created_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("cant scan query command %w", err)
	}

	return false, nil
}

func (d *MySqlDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec("INSERT INTO users (name, phone_number) VALUES (?, ?)", u.Name, u.PhoneNumber)
	if err != nil {
		return entity.User{}, fmt.Errorf("cant execute command %w", err)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}
