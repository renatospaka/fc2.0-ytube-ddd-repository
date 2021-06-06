package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/renatospaka/fc2.0-ytube-ddd-repository/entity"
)

type CategoryRepositorySQLite struct {
	db *sql.DB
}

func NewCategoryRepositorySQLite(db *sql.DB) *CategoryRepositorySQLite {
	return &CategoryRepositorySQLite{db: db}
}

func (c *CategoryRepositorySQLite) Get(id string) (entity.Category, error) {
	var category entity.Category
	stmt, err := c.db.Prepare("select id, name from categories where id=?")
	if err != nil {
		return entity.Category{}, err
	}
	//defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&category.ID, &category.Name)
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (c *CategoryRepositorySQLite) Create(category entity.Category) (entity.Category, error) {
	stmt, err := c.db.Prepare("insert into categories (id, name) values (?, ?)")
	if err != nil {
		return entity.Category{}, err
	}
	//defer stmt.Close()

	_, err = stmt.Exec(category.ID, category.Name)
	if err != nil {
		return entity.Category{}, err
	}

	err = stmt.Close()
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}