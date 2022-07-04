package auth

import (
	"github.com/Meysadesu/otakuread/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) RepositoryAuth {
	return &repository{db: db}
}

func (r *repository) Create(auth model.Auth) error {
	tx := r.db.Create(&auth)
	return tx.Error
}

func (r *repository) FindByUsername(username string) (model.Auth, error) {
	var auth model.Auth
	tx := r.db.Where("username = ?", username).First(&auth)
	return auth, tx.Error
}

func (r *repository) Count(username string) int64 {
	var auth model.Auth
	var count int64
	r.db.Model(&auth).Where("username = ?", username).Count(&count)
	return count
}
