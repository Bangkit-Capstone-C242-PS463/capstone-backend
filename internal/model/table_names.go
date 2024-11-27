package model

func (User) TableName() string {
	return "users"
}

func (History) TableName() string {
	return "history"
}