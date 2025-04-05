package database

type User struct{
    Id    uint 	   `gorm:"primaryKey;autoIncrement" json:"id"`  
    Ip    string   `gorm:"size:45;column:ip_address" json:"ip"`
}

type Polls struct {
    Id         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    UserId     uint   `gorm:"not null;index" json:"user_id"`
    MainHeader string `gorm:"size:255;not null" json:"main_header"`
    Type       string `gorm:"size:50;default:'general'" json:"type"`
    
    User       User   `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}