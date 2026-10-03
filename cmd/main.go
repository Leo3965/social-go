package main

import (
	"log"
	"time"

	"github.com/Leo3965/social/cmd/api"
	"github.com/Leo3965/social/internal/env"
	"github.com/Leo3965/social/internal/store/postgres"
)

const version = "0.0.1"

//	@title			Leo Social API
//	@description	This is a social server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	dbConfig := api.DbConfig{
		Addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
		MaxOpenConns: env.GetInteger("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: env.GetInteger("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", time.Minute*15),
	}

	cfg := api.Config{
		Addr:    env.GetString("ADDR", ":8080"),
		Db:      dbConfig,
		Env:     env.GetString("ENV", "development"),
		Version: version,
		ApiURL:  env.GetString("EXTERNAL_URL", "localhost:8080"),
	}

	db, err := postgres.New(dbConfig.Addr, dbConfig.MaxOpenConns, dbConfig.MaxIdleConns, dbConfig.MaxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	log.Println("database connection pool established")

	store := postgres.NewStorage(db)

	app := &api.Application{
		Config: cfg,
		Store:  store,
	}

	httpHandler := app.Mount()

	log.Fatal(app.Run(httpHandler))
}
