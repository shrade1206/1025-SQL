package main

// func Update() {
// 	// Update 只更新你選擇的字段
// 	// Updates 更新所有的字段 此時有兩種模式 一種為Map 一種為結構體 結構體0值不參與更新
// 	// Save 無論如何都更新 所有的內容 包含0值，使用上有點危險，請注意
// 	var pic []Pic

// 	Main_DB.Model(&pic).Where("Name = ? ", "test").Update("Name", "lol")

// 	dbRes := Main_DB.Where("Name = ? ", "test").Find(&pic)
// 	for k := range pic {
// 		pic[k].Name = "777"
// 	}
// 	dbRes.Save(pic)
// 	// 這種寫法無法更新 0值
// 	Main_DB.Where("Name = ? ", "test").Updates(Pic{Name: "", Picture: []byte("13")})
// 	// 這樣才能更新0 跟 空值
// 	Main_DB.Where("Name = ? ", "test").Updates(map[string]interface{}{"Name": "", "Picture": 0})
// }
