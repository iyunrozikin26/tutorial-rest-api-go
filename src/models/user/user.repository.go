// FYI! Repository file to handle call DB
package user

import (
	"fmt"

	"gorm.io/gorm"
)

// membuat interface untuk membuat kontrak method
type UserRepository interface {
	// nama, return method
	FindAll() []User
	FindOne(id int) User
	// pada Save, Update dan Delete, return valuenya adalah pointer agar tidak terjadinya duplikasi data yang mengakibatkan pemakaian memory yang besar
	Save(user User) (*User, error)
	Update(user User) (*User, error)
	Delete(user User) (*User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (ur *UserRepositoryImpl) FindAll() []User {
	var users []User
	_ = ur.db.Find(&users)
	fmt.Println(users, "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")
	return users
}

func (ur *UserRepositoryImpl) FindOne(id int) User {
	var user User
	_ = ur.db.First(&user, id)
	fmt.Println(user, "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")
	return user

}

func (ur *UserRepositoryImpl) Save(user User) (*User, error) {
	result := ur.db.Create(&user)
	fmt.Println(result, "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")

	if result != nil {
		return nil, result.Error
	}
	fmt.Println(&user, "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")
	return &user, nil
}

func (ur *UserRepositoryImpl) Update(user User) (*User, error) {
	result := ur.db.Save(&user)
	fmt.Println(result, "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")

	if result != nil {
		return nil, result.Error
	}
	fmt.Println(&user, "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")
	return &user, nil
}

func (ur *UserRepositoryImpl) Delete(user User) (*User, error) {
	result := ur.db.Delete(&user)
	fmt.Println(result, "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")
	if result != nil {
		return nil, result.Error
	}
	fmt.Println(&user, "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")
	return &user, nil
}
