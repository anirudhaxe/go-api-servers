package model

/*
These structures conform fully with the SQL schema
*/
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Todo struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
}
