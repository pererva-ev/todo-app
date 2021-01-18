package repository

import "github.com/jmoiron/sqlx"

type Task interface {
}

type Comment interface {
}

type Project interface {
}

type Column interface {
}

type Repository interface {
	Project
	Column
	Task
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
