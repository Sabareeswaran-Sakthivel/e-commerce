package models

type Customer struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Address   string `json:"address" db:"address"`
	Age       *int   `json:"age" db:"age"`
	Gender    string `json:"gender" db:"gender"`
	IsActive  bool   `json:"is_active" db:"is_active"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
