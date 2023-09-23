// FYI! Service file to handle logic
package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/iyunrozikin26/tutorial-rest-api-go.git/src/models/user/dto"
)

type UserService interface {
	GetAll() []User
	GetByID(id int) User
	Create(ctx *gin.Context) (*User, error) // ctx *gin.Context untuk menangkap params, query, dll dari input user
	Update(ctx *gin.Context) (*User, error)
	Delete(ctx *gin.Context) (*User, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

// constractor

func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (us *UserServiceImpl) GetAll() []User {
	return us.userRepository.FindAll()
}

func (us *UserServiceImpl) GetByID(id int) User {
	return us.userRepository.FindOne(id)
}

func (us *UserServiceImpl) Create(ctx *gin.Context) (*User, error) {
	var input dto.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}
	// mengisi struck dengan input dari client
	user := User{
		Name:  input.Name,
		Email: input.Email,
	}
	result, err := us.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (us *UserServiceImpl) Update(ctx *gin.Context) (*User, error) {
	id, _ := strconv.Atoi(ctx.Param("id")) // untuk mendapatkan id dan Atoi to mengconvert to ParseInt
	var input dto.UpdateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}
	// mengisi struck dengan input dari client
	user := User{
		ID:    int64(id),
		Name:  input.Name,
		Email: input.Email,
	}
	result, err := us.userRepository.Update(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (us *UserServiceImpl) Delete(ctx *gin.Context) (*User, error) {
	id, _ := strconv.Atoi(ctx.Param("id")) // untuk mendapatkan id dan Atoi to mengconvert to ParseInt 
	user := User{
		ID: int64(id),
	}
	result, err := us.userRepository.Delete(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
