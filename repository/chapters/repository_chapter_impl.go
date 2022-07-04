package chapters

import (
	"github.com/Meysadesu/otakuread/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepositoryChapter(db *gorm.DB) RepositoryChapters {
	return &repository{db: db}
}

func (r *repository) Create(chapters model.Chapter) error {
	tx := r.db.Create(&chapters).Error
	return tx
}

func (r *repository) FindByCH(chapter string) (model.Chapter, error) {
	var chapters model.Chapter
	tx := r.db.Where("ch = ?", chapter).First(&chapters).Error
	return chapters, tx
}
