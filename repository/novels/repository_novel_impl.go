package novels

import (
	"github.com/Meysadesu/otakuread/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepositoryNovel(db *gorm.DB) RepositoryNovels {
	return &repository{db: db}
}

func (r *repository) Create(novel model.Novel) error {
	tx := r.db.Create(&novel)
	return tx.Error
}

func (r *repository) Find() ([]model.Novel, error) {
	var novel []model.Novel
	tx := r.db.Find(&novel)
	return novel, tx.Error
}

func (r *repository) FindByTitle(title string) (model.Novel, error) {
	var novel model.Novel
	var volume []model.Volume
	tx := r.db.Where("title LIKE ?", "%"+title+"%").First(&novel)
	r.db.Where("title_novel = ?", novel.Title).Find(&volume)
	novel.Volume = volume
	return novel, tx.Error
}

func (r *repository) Delete(id string) error {
	var novel model.Novel
	tx := r.db.Where("id = ?", id).Delete(&novel)
	return tx.Error
}

func (r *repository) Update(novels model.Novel) error {
	tx := r.db.Where("id = ?", novels.ID).Updates(&novels)
	return tx.Error
}
