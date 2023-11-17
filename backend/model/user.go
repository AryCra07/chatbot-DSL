package model

// User of a store
type User struct {
	Id       string  `gorm:"primary_key;" json:"id"`
	Name     string  `gorm:"type:varchar(20);not null;unique;" json:"name"`
	Password string  `gorm:"type:varchar(64);not null;" json:"password"`
	State    int32   `gorm:"type:int32;" json:"state"`
	Balance  float32 `gorm:"type:float32;" json:"balance"`
	Bill     float32 `gorm:"type:float32;" json:"bill"`
}
