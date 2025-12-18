package dummy

import (
	"fmt"

	"challenge/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

// To manage more than one user object
var users []User

// Add one user
func UserAdd(u *User) (err error) {

	if err = db.Conn.Create(u).Error; err != nil {
		return err
	} else {
		return nil
	}

}

// Delete one user
func UserDel(id string) (err error) {
	result := db.Conn.Delete(&User{}, id)

	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

// Update one user
func UserUpdate() {
	fmt.Println("Not implemented yet.")
}

// Get one user
func UserGet(u *User, id string) (err error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE ID = %v", id)
	err = db.Conn.Raw(query).Scan(u).Error
	return err
}

// List all users
func UserList() ([]User, int64, error) {

	results := db.Conn.Find(&users)
	return users, results.RowsAffected, results.Error

}

// Find one user
func UserFind(u User) {
	fmt.Printf("You are looking for user: %v, but that function is not implemented yet.\n", u)
}
