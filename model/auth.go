package model

type Auth struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"size:100;not null;unique" json:"username" validate:"required"`
	Email    string `gorm:"size:100;not null;unique" json:"email" validate:"required,email"`
	Role     string `gorm:"size:20;not null;default:'user'" json:"role"`
	Password string `gorm:"not null" json:"password" validate:"required"`
}

type UserResponse struct {
	ID       int
	Username string
	Email    string
	Password string
}

type AdminResponse struct {
	ID       int
	Username string
	Email    string
	Password string
	Token    string
}
