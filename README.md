# tutorial-rest-api-go
1. Melakukan `go mod init example.com/greetings`
2. Melakukan installasi framework gin `go get -u github.com/gin-gonic/gin`
3. Menginstall go orm : `go get -u gorm.io/gorm`
4. Menginstall driver database : `go get -u gorm.io/driver/mysql` (tergantung dari db yang ingin digunakan)
5. Selanjutnya folder src yang isinya adalah folder config (untuk configurasi database) dan di dalam config folder terdapat db.go. Setup dari db.go :
from `https://gorm.io/docs/connecting_to_the_database.html#MySQL`
    ```go
    package config
    import (
        "gorm.io/driver/mysql"
        "gorm.io/gorm"
    )

    func DB() *gorm.DB {
        host := "localhost"
        port := "3306"
        dbname := "go-tutorial"
        username := "root"
        password := "root"

        dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
        db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

        if err != nil {
            panic("Tidak dapat terkoneksi ke database")
        }
        return db
    }
    ```
6. buatlah database sesuai dengan dbname
7. siapkan file main.go yang setupnya sebagai berikut:
    ```go

    ```
8. Setup models
    - user.model.go
        ```go
        // untuk membuat cetakan object dari user
        package user
        type User struct {
            ID    int64  `json:"id" gorm:"primaryKey;auto_increament:true;index"`
            Name  string `json:"name" gorm:"type:varchar(255)"`
            Email string `json:"email" gorm:"type:varchar(255)"`
        }        
        ``` 
    - user.repository.go
        ```go
        // FYI! Repository file to handle call DB
        package user

        import "gorm.io/gorm"

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
            return users
        }

        func (ur *UserRepositoryImpl) FindOne(id int) User {
            var user User
            _ = ur.db.First(&user, id)
            return user

        }

        func (ur *UserRepositoryImpl) Save(user User) (*User, error) {
            result := ur.db.Create(&user)
            if result != nil {
                return nil, result.Error
            }
            return &user, nil
        }

        func (ur *UserRepositoryImpl) Update(user User) (*User, error) {
            result := ur.db.Save(&user)
            if result != nil {
                return nil, result.Error
            }
            return &user, nil
        }

        func (ur *UserRepositoryImpl) Delete(user User) (*User, error) {
            result := ur.db.Delete(&user)
            if result != nil {
                return nil, result.Error
            }
            return &user, nil
        }
        ``` 
    - user.service.go
        ```go
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
            result, err := us.userRepository.Save(user)
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
        ```
    - user.controller.go
        ```go
        // FYI! Controller file to handle request and response client
        package user

        import (
            "net/http"
            "strconv"

            "github.com/gin-gonic/gin"
        )

        type UserController struct {
            userService UserService
            ctx         *gin.Context
        }

        func NewUserController(userService UserService, ctx *gin.Context) UserController {
            return UserController{userService, ctx}
        }

        func (uc *UserController) Index(ctx *gin.Context) {
            data := uc.userService.GetAll()
            ctx.JSON(http.StatusOK, gin.H{
                "status": "OK",
                "data":   data,
            })
        }

        func (uc *UserController) GetById(ctx *gin.Context) {
            id, _ := strconv.Atoi(ctx.Param("id"))
            data := uc.userService.GetByID(id)
            ctx.JSON(http.StatusOK, gin.H{
                "status": "OK",
                "data":   data,
            })
        }

        func (uc *UserController) Create(ctx *gin.Context) {
            data, err := uc.userService.Create(ctx)
            if err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{
                    "status": "Error",
                    "data":   err,
                })
                ctx.Abort()
                return 
            }
            ctx.JSON(http.StatusOK, gin.H{
                "status": "OK",
                "data":   data,
            })
        }
        func (uc *UserController) Delete(ctx *gin.Context) {
            data, err := uc.userService.Delete(ctx)
            if err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{
                    "status": "Error",
                    "data":   err,
                })
                ctx.Abort()
                return 
            }
            ctx.JSON(http.StatusOK, gin.H{
                "status": "OK",
                "data":   data,
            })
        }
        func (uc *UserController) Update(ctx *gin.Context) {
            data, err := uc.userService.Update(ctx)
            if err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{
                    "status": "Error",
                    "data":   err,
                })
                ctx.Abort()
                return 
            }
            ctx.JSON(http.StatusOK, gin.H{
                "status": "OK",
                "data":   data,
            })
        }
        ```
9. Setup 

