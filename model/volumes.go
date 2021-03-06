package model

type Volume struct {
	ID          int       `gorm:"primaryKey;autoIncrement;not null"`
	TitleNovel  string    `gorm:"size:255; not null" json:"title_novel"`
	WhatsVolume string    `gorm:"size:255; not null" json:"volume"`
	Novel       Novel     `gorm:"foreignKey:TitleNovel;references:Title"`
	Chapter     []Chapter `gorm:"foreignKey:IDVolume;references:ID"`
}

type VolumeResponse struct {
	ID          int
	TitleNovel  string
	WhatsVolume string
	Chapter     []Chapter
}
