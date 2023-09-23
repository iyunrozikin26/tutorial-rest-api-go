package user

// untuk membuat cetakan object dari user
type User struct {
	ID    int64  `json:"id" gorm:"primaryKey;auto_increament:true;index"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Email string `json:"email" gorm:"type:varchar(255)"`
}
