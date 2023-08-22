package models

type Photos struct {
	ID       int16  `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:text" json:"caption"`
	PhotoUrl string `gorm:"type:varchar" json:"photo_url"`
	UserID   int    `gorm:"index" json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"user"`
}
