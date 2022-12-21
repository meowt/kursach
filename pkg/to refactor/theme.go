package to_refactor

//
//type Theme struct {
//	ID          int
//	Name        string
//	Path        string
//	ReleaseDate pgtype.Date
//	CreatorName string
//	Description sql.NullString
//}
//
//// Save function inserts theme data into Themes table
//func (t *Theme) Save() error {
//	sqlQuery := fmt.Sprintf("INSERT INTO themes (name, path, creator_name, description) VALUES ('%s, %s, %s, %s');", t.Name, t.Path, t.CreatorName, t.Description.String)
//	_, err := PostgresClient.Exec(sqlQuery)
//	return err
//}
//
//// GetByID returns data about
//func (t *Theme) GetByID(id string) error {
//	err := PostgresClient.QueryRow(fmt.Sprintf("SELECT * FROM themes WHERE id = '%s' ;", id)).Scan(
//		&t.Path,
//		&t.ReleaseDate,
//		&t.CreatorName,
//		&t.Description,
//		&t.ID,
//		&t.Name)
//	return err
//}
//
//func (t *Theme) GetWithLimit(limit int) error {
//	err := PostgresClient.QueryRow(fmt.Sprintf("SELECT * FROM themes WHERE id = '%v' ;", limit)).Scan(
//		&t.Path,
//		&t.ReleaseDate,
//		&t.CreatorName,
//		&t.Description,
//		&t.ID,
//		&t.Name)
//	return err
//}
//
//// GetLastThemeID returns id of last theme stored in database
//func GetLastThemeID() int {
//	var id int
//	_ = PostgresClient.QueryRow(fmt.Sprintf("SELECT max(id) FROM themes;")).Scan(&id)
//	return id
//}
//
//// GetCreatorsThemes gets data about 4 themes of
//func GetCreatorsThemes(creator string) ([]Theme, error) {
//
//	var themes []Theme
//
//	rows, err := PostgresClient.Query(fmt.Sprintf("SELECT * FROM themes WHERE creator_name = '%s' LIMIT 4;", creator))
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	for rows.Next() {
//		var theme Theme
//		err = rows.Scan(
//			&theme.Path,
//			&theme.ReleaseDate,
//			&theme.CreatorName,
//			&theme.Description,
//			&theme.ID,
//			&theme.Name)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		themes = append(themes, theme)
//	}
//
//	return themes, err
//}
//
//// GetLastThemes f
//func GetLastThemes() ([]Theme, error) {
//	var themes []Theme
//
//	rows, err := PostgresClient.Query("SELECT * FROM themes ORDER BY release_date LIMIT 4;")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	for rows.Next() {
//		var theme Theme
//		err = rows.Scan(
//			&theme.Path,
//			&theme.ReleaseDate,
//			&theme.CreatorName,
//			&theme.Description,
//			&theme.ID,
//			&theme.Name)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		themes = append(themes, theme)
//	}
//
//	return themes, err
//}
