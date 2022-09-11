package server

import (
	_ "github.com/Ajay-Jagtap382/Library-Management-System/app"
	// "github.com/Ajay-Jagtap382/Library-Management-System/db"
	// "github.com/Ajay-Jagtap382/Library-Management-System/users"
)

type dependencies struct {
	// CategoryService users.Service
}

func initDependencies() (dependencies, error) {
	// appDB := app.GetDB()
	// logger := app.GetLogger()
	// dbStore := db.NewStorer(appDB)

	// categoryService := users.NewService(dbStore, logger)

	return dependencies{
		// CategoryService: categoryService,
	}, nil
}
