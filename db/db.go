package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetDB() (*gorm.DB, error) {
	db, err := NewDB(&Option{
		Host:        "127.0.0.1",
		Port:        3306,
		User:        "root",
		Password:    "Q1w@e3r4",
		Database:    "callonline",
		TablePrefix: "c_",
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Option struct {
	Host        string
	Port        int
	User        string
	Password    string
	Database    string
	TablePrefix string
}

func NewDB(option *Option) (*gorm.DB, error) {
	address := fmt.Sprintf("%s:%d", option.Host, option.Port)
	db, err := dbConnect(
		option.User,
		option.Password,
		address, option.Database,
		option.TablePrefix)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func dbConnect(user, pass, addr, dbname, tableprefix string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pass,
		addr,
		dbname,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tableprefix,
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
