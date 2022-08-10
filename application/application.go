package application

import (
	"fmt"
	"orinz/application/drivers/database"
	"orinz/application/drivers/database/repository"
	usecase "orinz/application/usecase/class"
)

type Engine struct {
	createClassUseCase usecase.CreateClassUseCase
}

func New() *Engine {
	engine := &Engine{}
	engine.loadDependencies()
	return engine
}

func (e *Engine) loadDependencies() {
	db := database.ConnectSQL()
	fmt.Println("database connection successfully")

	classRepository := repository.NewClassRepository(db)
	e.createClassUseCase = usecase.NewCreateClass(classRepository)
}
