package schema

import (
	"template_clean_architecture/utils/helpers"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account        Account   `gorm:"foreignKey:user_id"`
	Password       string    `gorm:"column:password;not null"`
	Email          string    `gorm:"column:email;unique;not null"`
	LastAccessedAt time.Time `gorm:"column:last_accessed_at"`
}

// compare password
func (u *User) ComparePassword(password string) bool {
	return helpers.CheckPasswordHash(password, u.Password)
}
