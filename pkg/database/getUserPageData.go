package database

import "fmt"

func GetUserPageData(d string) (User, error) {
	var pageOwner User

	e := db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE username = '%s' ;", d)).Scan(&pageOwner.Username, &pageOwner.Email, &pageOwner.Password, &pageOwner.Description, &pageOwner.ID)
	if e != nil {
		fmt.Println(e.Error())
	}

	return pageOwner, e
}

func GetThemesData(d string) ([]Theme, error) {
	var themes []Theme

	rows, e := db.Query(fmt.Sprintf("SELECT * FROM themes WHERE creator_name = '%s' LIMIT 4;", d))
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
