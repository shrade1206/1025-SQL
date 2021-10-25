package main

import "log"

// 設定Table
type Pic struct {
	Id      int `gorm:"primary_key"`
	Name    string
	Picture []byte `gorm:"type:blob"`
}

func TestCreate() {
	// AutoMigrate 隨時更新 Table狀況
	err := Main_DB.AutoMigrate(&Pic{})
	if err != nil {
		log.Printf("Table Error : %s", err.Error())
	}
}
