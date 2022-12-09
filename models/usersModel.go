package models

// User Type -> userss table
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// add user
func AddUser(user User) (id int, err error) {
	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return user.ID, nil
	}
}
