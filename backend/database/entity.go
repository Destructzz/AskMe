package database

type User struct{
    Id    uint 	   `gorm:"primaryKey;autoIncrement" json:"id"`  
    Ip    string   `gorm:"size:45;column:ip_address" json:"ip"`
}

type Poll struct {
    Id         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    UserId     uint   `gorm:"not null;index" json:"user_id"`
    MainHeader string `gorm:"size:255;not null" json:"main_header"`
    Type       string `gorm:"size:50;default:'general'" json:"type"`

    User       *User   `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

type Card struct{
    Id uint `gorm:"primaryKey;autoIncrement" json:"id"`
    Header string `gorm:"size:255;not null" json:"header"`
    PollId uint `gorm:"not null;index;" json:"poll_id"`

    Poll *Poll `gorm:"foreignKey:PollId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"poll"`
}

type Card_Field struct {
    Id         uint   `gorm:"primaryKeyм;autoIncrement" json:"id"`
    CardId    uint   `gorm:"not null;index" json:"card_id"`
    Text     string   `gorm:"size:255;not null" json:"header"`

    Card *Card        `gorm:"foreignKey:CardId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"card"`
}



