package volumes

import (
	"github.com/Meysadesu/otakuread/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRespositoryVolumes(db *gorm.DB) RepositoryVolumes {
	return &repository{db: db}
}

func (r *repository) Create(volume model.Volume) error {
	tx := r.db.Create(&volume)
	return tx.Error
}

func (r *repository) FindAll() ([]model.Volume, error) {
	var volumes []model.Volume
	tx := r.db.Find(&volumes)
	return volumes, tx.Error
}

func (r *repository) FindByID(id string) (model.Volume, error) {
	var volume model.Volume
	var chapter []model.Chapter
	tx := r.db.Where("id = ?", id).First(&volume)
	r.db.Where("id_volume =?", volume.ID).Find(&chapter)
	volume.Chapter = chapter
	return volume, tx.Error
}
