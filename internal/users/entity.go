package users

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (u User) TableName() string {
	return "public.user"
}
