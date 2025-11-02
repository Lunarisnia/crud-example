package userentities

type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (u User) TableName() string {
	return "public.user"
}
