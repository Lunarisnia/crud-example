package userentities

type User struct {
	ID   uint   `gorm:"primaryKey" json:"id" redis:"id"`
	Name string `json:"name" redis:"name"`
}

func (u User) TableName() string {
	return "public.user"
}
