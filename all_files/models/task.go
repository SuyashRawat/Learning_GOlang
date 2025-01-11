package models

type Task struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Status      string `db:"status"`
	ProjectID   int    `db:"project_id"`
	AssignedTo  int    `db:"assigned_to"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
