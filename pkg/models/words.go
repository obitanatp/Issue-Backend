package models

type Word struct {
	ID     uint         `gorm:"primaryKey" json:"id"`
	Title  string       `json:"title"`
	Answer []WordAnswer `gorm:"foreignKey:idques;references:id;constraint:OnInsert:CASCADE,OnUpdate:CASCADE,OnDelete:CASCADE;" json:"answer"`
}

type WordAnswer struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	IDques      uint   `json:"idques"`
	Title       string `json:"titleAnswer"`
	Constraint1 string `json:"constraint1"`
	Constraint2 string `json:"constraint2"`
}
