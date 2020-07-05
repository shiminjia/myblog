package dbops

import "log"

type UserLogin struct {
	Id       int
	Username string
	Email    string
	Password string
	Time     *Time
}



func SelectUserByEmail(email string) UserLogin {
	var dao UserLogin
	err := DB.QueryRow("select id, email, password, username from users where email = ?", email).Scan(&dao.Id, &dao.Email, &dao.Password, &dao.Username)
	//&dao.Time.Created_at, &dao.Time.Updated_at, &dao.Time.Deleted_at
	if err != nil {
		log.Println(err.Error())
		return dao
	}
	return dao
}

