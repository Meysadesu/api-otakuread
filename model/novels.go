package model

type Novel struct {
	ID          int      `gorm:"primaryKey;autoIncrement;unique;not null" json:"id"`
	Title       string   `gorm:"size:255;unique;not null" json:"title" validate:"required"`
	Alternative string   `gorm:"size:255;" json:"alternative"`
	Cover       string   `gorm:"size:255;not null;default:'notfound.jpg'" json:"cover"`
	Type        string   `gorm:"size:100;not null" json:"type"`
	Author      string   `gorm:"size:255" json:"author"`
	Status      string   `gorm:"size:10;default:'Ongoing'" json:"status"`
	Ranting     float64  `gorm:"default:1.0" json:"ranting"`
	Genre       string   `gorm:"not null" json:"genre"`
	Description string   `json:"description"`
	Volume      []Volume `gorm:"foreignKey:TitleNovel;references:Title"`
}

type NovelResponse struct {
	ID          int
	Title       string
	Alternative string
	Cover       string
	Type        string
	Author      string
	Status      string
	Ranting     float64
	Genre       string
	Description string
}
