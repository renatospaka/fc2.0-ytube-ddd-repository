package main

import (
	"database/sql"
	"fmt"

	_ "github.com/renatospaka/fc2.0-ytube-ddd-repository/entity"
	"github.com/renatospaka/fc2.0-ytube-ddd-repository/repository"
	"github.com/renatospaka/fc2.0-ytube-ddd-repository/service"
)

func main() {
	// in-memory repository
	// db := repository.CategoriesMemoryDb{Categories: []entity.Category{}}
	// repositoryMemory := repository.NewCategoryRepositoryMemory(db)
	// serv := service.NewCategoryService(repositoryMemory)

	// SQLite repository
	db, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySQLite := repository.NewCategoryRepositorySQLite(db)
	serv := service.NewCategoryService(repositorySQLite)

	cat, _ := serv.Create("Minha Categoria")
	fmt.Println(cat)
}