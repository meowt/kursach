package database

import "fmt"

func GetCatalogue(s string) ([]Theme, error) {
	var themes []Theme
	var sql string
	if s == "LAST" {
		sql = "SELECT * FROM themes ORDER BY release_date LIMIT 4;"
	} else if s == "MEOWT" {
		sql = "SELECT * FROM themes WHERE creator_name = 'meowt' LIMIT 4;"
	} else {
		sql = "SELECT * FROM themes LIMIT 4;"
	}

	rows, e := db.Query(sql)
	if e != nil {
		fmt.Println(e.Error())
	}
	for rows.Next() {
		var theme Theme
		e = rows.Scan(
			&theme.Path,
			&theme.ReleaseDate,
			&theme.CreatorName,
			&theme.Description,
			&theme.ID,
			&theme.Name)
		if e != nil {
			fmt.Println(e.Error())
		}
		themes = append(themes, theme)
	}

	return themes, e
}
