package models

import (
	"database/sql"
)



type User struct{
	Id string
	UserName string
	Password string
	NickName string
	RegistTime string
	LastTimeLogin sql.NullString
	NewLoginTime sql.NullString
	Bak sql.NullString
	Online sql.NullString
	CreateTime sql.NullString
	Creator sql.NullString
	UpdateTime sql.NullString
	Updator sql.NullString
}
