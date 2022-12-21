package database

//type User struct {
//	ID          int            `json:"id"`
//	Username    string         `json:"username"`
//	Email       string         `json:"email"`
//	Password    string         `json:"password"`
//	Description sql.NullString `json:"description"`
//}
//
//func (user *User) GetPageData(d string) error {
//
//	err := PostgresClient.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE username = '%s' ;", d)).Scan(&user.Username, &user.Email, &user.Password, &user.Description, &user.ID)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	return err
//}
//
//func (user *User) LoginRequest(email, password string) string {
//	var passwordCheck string
//
//	sqlQuery := fmt.Sprintf("SELECT password FROM users WHERE email = '%s';", email)
//
//	err := PostgresClient.QueryRow(sqlQuery).Scan(&passwordCheck)
//	if err != nil && err != sql.ErrNoRows {
//		return "Произошла ошибка при авторизации"
//	}
//	if err == sql.ErrNoRows {
//		return "Пользователь с такой почтой не зарегистрирован"
//	}
//	if service.ComparePassword(password, passwordCheck) {
//		sqlQuery = fmt.Sprintf("SELECT * FROM users WHERE email = '%v';", email)
//		err = PostgresClient.QueryRow(sqlQuery).Scan(&user.Username, &user.Email, &user.Password, &user.Description, &user.ID)
//		return ""
//	} else {
//		return "Неправильный пароль"
//	}
//}
//
//func (user *User) RegRequest() string {
//	//Checking of existing users with this email
//
//	var scanStruct User
//	row := PostgresClient.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE email = '%s' LIMIT 1;", user.Email))
//	e := row.Scan(&scanStruct.Username, &scanStruct.Email, &scanStruct.Password, &scanStruct.Description, &scanStruct.ID)
//
//	if e != nil && e != sql.ErrNoRows {
//		return "error"
//	}
//
//	userID := strconv.Itoa(scanStruct.ID)
//
//	if userID != "0" {
//		return "Пользователь с такой почтой уже зарегистрирован"
//	}
//
//	//Checking of existing users with this email
//	row = PostgresClient.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE username = '%s' LIMIT 1;", user.Username))
//	e = row.Scan(&scanStruct.Username, &scanStruct.Email, &scanStruct.Password, &scanStruct.Description, &scanStruct.ID)
//	if e != nil && e != sql.ErrNoRows {
//		return "error"
//	}
//
//	userName := scanStruct.Username
//
//	if userName != "" {
//		return "Пользователь с таким именем уже существует"
//	}
//
//	//Hashing password
//	hashedPassword, e := service.HashPassword(user.Password)
//	if e != nil {
//		return "error"
//	}
//
//	//Inserting new user into PostgresClient
//	sqlQuery := fmt.Sprintf("INSERT INTO users (username, email, password) VALUES ('%s', '%s', '%s');", user.Username, user.Email, hashedPassword)
//	_, e = PostgresClient.Exec(sqlQuery)
//	if e != nil {
//		return "error"
//	}
//	return ""
//}
