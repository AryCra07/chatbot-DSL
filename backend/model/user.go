package model

type User struct {
	Id       string `gorm:"primary_key" json:"id"`
	Name     string `gorm:"type:varchar(20);not null;unique" json:"user"`
	Password string `gorm:"type:varchar(64);not null;" json:"password"`
	Mode     string `gorm:"type:varchar"`
}

func (User) TableName() string {
	return "user"
}
