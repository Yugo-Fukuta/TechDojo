package model

type User struct {
	ID       int    `json:"id"`
	Name 	 string `json:"name"`
	Token 	 string `json:"token"`
}

func CreateUser(user *User)  {
	db := Init()
	db.Create(user)
}

func GetUser(user *User) *User {
	db := Init()
	db.Find(user)
	return user
}

func UpdateUSer(user *User) {
	db := Init()
	db.Save(user)
}