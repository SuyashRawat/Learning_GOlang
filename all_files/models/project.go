package models

type Project struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	OwnerID     int    `db:"owner_id"`
	CreatedAt   string `db:"created_at"`
}
