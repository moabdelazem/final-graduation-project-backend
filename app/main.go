package main

// Application EntryPoint
func main() {
	// Create Global Configurations
	appConfig := Config{
		addr: ":8080",
	}
	// Create New Application Instance
	app := Application{
		config: appConfig,
	}
	// Start Server
	app.StartServer(app.Mount())
}
