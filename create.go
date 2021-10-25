package main

// import "log"

// func Create() {
// 	// 新增一筆資料
// 	err := Main_DB.Create(&Pic{Name: "GGG", Picture: "Y"})
// 	if err != nil {
// 		log.Printf("Created Error : %s", err.Error)
// 	}
// 	// 指定要新增值的欄位
// 	Main_DB.Omit("Name").Create(&Pic{Name: "GGG", Picture: "Y"})

// 	// 大量資料一次新增
// 	Main_DB.Create(&[]Pic{
// 		{Name: "GGG", Picture: "Y"},
// 		{Name: "aaa", Picture: "N"},
// 		{Name: "ccc", Picture: "A"},
// 		{Name: "vvv", Picture: "E"},
// 		{Name: "rrr", Picture: "F"},
// 		{Name: "ttt", Picture: "B"},
// 	})
// }
