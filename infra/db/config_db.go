package db

import "fmt"

//TODO criar struct para isto

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "db"

const Port = "5432"

const Password = "123456"

const DbName = "bank"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
