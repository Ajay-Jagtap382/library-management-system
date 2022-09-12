package server

import (
	"github.com/Ajay-Jagtap382/library-management-system/app"
	"github.com/Ajay-Jagtap382/library-management-system/db"
	"github.com/Ajay-Jagtap382/library-management-system/users"
)

type dependencies struct {
	UserService users.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := users.NewService(dbStore, logger)

	return dependencies{
		UserService: userService,
	}, nil
}
