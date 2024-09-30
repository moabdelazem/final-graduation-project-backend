package main

import (
	"github.com/joho/godotenv"
	"github.com/moabdelazem/final-graduation-project-backend/internals/db"
	"github.com/moabdelazem/final-graduation-project-backend/internals/env"
)

// Application EntryPoint
func main() {
	godotenv.Load()

	// Create Global Configurations
	appConfig := Config{
		addr: ":8080",
		db: dbConfig{
			addr: env.GetEnvString("DB_ADDR", ""),
		},
	}

	// Create New DB Connection
	_, err := db.NewDB(appConfig.db.addr)
	if err != nil {
		panic(err)
	}

	// Create New Application Instance
	app := Application{
		config: appConfig,
	}
	// Start Server
	app.StartServer(app.Mount())
}
