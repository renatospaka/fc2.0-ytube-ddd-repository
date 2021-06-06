package main

import (
	"fmt"

	"github.com/renatospaka/fc2.0-ytube-ddd-repository/entity"
	"github.com/renatospaka/fc2.0-ytube-ddd-repository/repository"
	"github.com/renatospaka/fc2.0-ytube-ddd-repository/service"
)

func main() {
	db := repository.CategoriesMemoryDb{Categories: []entity.Category{}}
	repositoryMemory := repository.NewCategoryRepositoryMemory(db)

	serv := service.NewCategoryService(repositoryMemory)
	cat, _ := serv.Create("Minha Categoria")

	fmt.Println(cat)
}