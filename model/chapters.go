package model

type Chapter struct {
	ID       int    `gorm:"primaryKey;autoIncrement;not null"`
	IDVolume int    `gorm:"not null" json:"id_volume"`
	Volume   Volume `gorm:"foreignKey:IDVolume;references:ID"`
	Ch       string `gorm:"not null;size:255;" json:"chapter"`
	Header   string `gorm:"not null" json:"header"`
	Ctx      string `json:"content"`
}

type ChapterResponse struct {
	ID       int
	IDVolume int
	Ch       string
	Header   string
	Ctx      string
}
