package mysql

import entity "suggestApp/enity"

func (d DB) IsUniquePhoneNumber(phoneNumber string) (bool, error) {}
func (d DB) Register(u entity.User) (entity.User, error)          {}
