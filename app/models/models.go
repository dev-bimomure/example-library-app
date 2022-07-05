package models

import "database/sql"

type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Book struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Tanggal   string `json:"tanggal"`
}
