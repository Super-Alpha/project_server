package db

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

type Enterprise struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func TestNewDB(t *testing.T) {
	db, err := GetDB()
	if err != nil {
		fmt.Println(err)
	}
	item := make([]Enterprise, 0)

	db.Find(&item)

	fmt.Println(item)
}
