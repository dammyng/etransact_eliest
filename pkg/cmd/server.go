package cmd

import (
	"eliest/internals/db"
	"eliest/internals/implementation"
	"eliest/models"
	"eliest/pkg"
	"log"
)


var port = "0.0.0.0:5001" 

// RunServer -
func RunServer() error {
	var eliest = pkg.Eliest{}

	
	//The Db 
	DBConfig := db.InitConfig()
	sqlDb := implementation.NewSqlLayer((db.Config(&DBConfig)))
	sqlDb.Session.AutoMigrate(models.Account{})
	err := eliest.InitializeDb(sqlDb)
	if err != nil {
		log.Printf("RunServer() - Failed to initialize db with error %v", err)
		return err
	}

	//The Router
	eliest.SetRoutes(eliest.Db )

	err = eliest.StartHttp(port)
	return err
}
