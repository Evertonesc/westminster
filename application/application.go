package application

import (
	"fmt"
	"orinz/application/drivers/database"
)

type Engine struct {
}

func New() *Engine {
	engine := &Engine{}
	engine.loadDependencies()
	return engine
}

func (e *Engine) loadDependencies() {
	db := database.ConnectSQL()
	fmt.Println("database connection successfully")
	fmt.Println(db)
}
