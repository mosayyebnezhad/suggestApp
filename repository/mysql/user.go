package mysql

import (
	"database/sql"
	"fmt"
	entity "suggestApp/enity"
)

func (d *MySqlDB) IsUniquePhoneNumber(phoneNumber string) (bool, error) {

	row := d.db.QueryRow("select * from users where phone_number= ?", phoneNumber)

	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("cant scan query command %w", err)
	}

	return false, nil
}

func (d *MySqlDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec("INSERT INTO users (name, phone_number, password) VALUES (?, ?, ?)", u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("cant execute command %w", err)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}

func (d *MySqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	row := d.db.QueryRow("select * from users where phone_number= ?", phoneNumber)

	user, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, fmt.Errorf("cant scan query command %w", err)
	}

	return user, true, nil
}

func (d *MySqlDB) GetUserByID(id uint) (entity.User, error) {

	row := d.db.QueryRow("select * from users where id= ?", id)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("record not found %w", err)
		}

		return entity.User{}, fmt.Errorf("cant scan query command %w", err)
	}

	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	var (
		user       entity.User
		created_at []uint8
	)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &created_at)

	return user, err
}
