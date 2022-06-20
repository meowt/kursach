package database

import "fmt"

func GetTheme(d string) (Theme, error) {
	var theme Theme

	e := db.QueryRow(fmt.Sprintf("SELECT * FROM themes WHERE id = '%s' ;", d)).Scan(
		&theme.Path,
		&theme.ReleaseDate,
		&theme.CreatorName,
		&theme.Description,
		&theme.ID,
		&theme.Name)
	if e != nil {
		fmt.Println(e.Error())
	}

	return theme, e
}

func GetLastThemeId() int {
	var id int
	_ = db.QueryRow(fmt.Sprintf("SELECT id FROM themes ORDER BY id DESC LIMIT 1;")).Scan(&id)
	return id
}
