package database

func SaveTheme(t Theme) error {
	//inserting new user into db
	sqlQuery := "INSERT INTO themes (name, path, creator_name, description) VALUES ('" + t.Name + "', '" + t.Path + "', '" + t.CreatorName + "', '" + t.Description.String + "');"
	_, e := db.Exec(sqlQuery)
	return e
}
