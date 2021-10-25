package main

import (
	"log"
)

func Find() {
	// 跟下面make()是一樣的
	// result := map[string]interface{}{}
	// result := make(map[string]interface{})
	// Main_DB.Model(&Pic{}).First(&result)
	// log.Println(result)

	//宣告結構體也可以查詢
	var Aa Pic
	Main_DB.Model(&Pic{}).First(&Aa)
	// Main_DB.Last(&Aa, map[string]interface{}{"Name": "Aa"})
	// Main_DB.Last(&Aa, Pic{Name: "Aa"})

	// // SELECT * FROM `users` LIMIT 5 OFFSET 10
	// Main_DB.Offset(10).Limit(5).Find(&Aa)

	// // pk檢索 後面只能接數字
	// // err := Main_DB.Model(&Pic{}).First(&Aa,1)
	// Main_DB.Model(&Pic{}).Last(&Aa)
	// log.Println(Aa)
	// // 查錯誤
	// fmt.Println(errors.Is(err.Error, gorm.ErrRecordNotFound))

	// // 條件查詢
	// Main_DB.Where("Name = ? ", "test").First(&Aa)

	// // <> 是不等於
	// Main_DB.Where("Name <> ? ", "test").Find(&Aa)

	// // 查詢 Name 跟 Picture 或是 Name 符合的條件
	// Main_DB.Where("Name = ? AND Picture =? ", "test", "TTT").Or("Name = ? ", "Hello").First(&Aa)

	// // id = 2 則不符合查詢條件
	// Main_DB.Where("Name = ? AND Picture =? ", "test", "TTT").Not("id = ?", 2).First(&Aa)

	// Main_DB.Where(map[string]interface{}{
	// 	"Name = ? ": "test",
	// }).First(&Aa)

	//語法完，直接印出結構體即可拿到資料
	log.Println(Aa)
}
