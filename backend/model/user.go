package model

type User struct {
	Id       string `gorm:"primary_key" json:"id"`
	Name     string `gorm:"type:varchar(20);not null;unique" json:"user"`
	Password string `gorm:"type:varchar(20);not null;" json:"password"`
}

func (User) TableName() string {
	return "user"
}
