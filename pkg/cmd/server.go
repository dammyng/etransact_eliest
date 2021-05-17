package cmd

import (
	"eliest/internals/db"
	"eliest/internals/implementation"
	"eliest/logger/implement"
	"eliest/models"
	"eliest/myredis"
	"eliest/pkg"
	"log"
	"os"
)


var port = "0.0.0.0:5001" 

// RunServer -
func RunServer() error {
	var eliest = pkg.Eliest{}

	
	//The Db 
	DBConfig := db.InitConfig()
	ADBConfig := db.InitAConfig()

	sqlDb := implementation.NewSqlLayer((db.Config(&DBConfig)), (db.Config(&ADBConfig)))
	sqlDb.Session.AutoMigrate(models.Account{}, models.Winnings{}, models.Vouchers{})
	err := eliest.InitializeDb(sqlDb)
	if err != nil {
		log.Printf("RunServer() - Failed to initialize db with error %v", err)
		return err
	}

	redisClient := myredis.NewRedisClient(os.Getenv("Redis_Host"), os.Getenv("RedisPass"))

	newLogger := implement.NewGamesFileSystem()
	//The Router
	eliest.SetRoutes(eliest.Db , newLogger, redisClient)

	err = eliest.StartHttp(port)
	return err
}
