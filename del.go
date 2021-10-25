package main

func del() {
	var pic []Pic
	Main_DB.Where("Id = ? ", 3).Delete(&pic)
	//
	// Main_DB.Unscoped().Where("Name = ?", "Test").Delete(&pic)

}
