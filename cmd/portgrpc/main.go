package portgrpc

func main() {
	app, err := BuildDIForApp()
	if err != nil {
		// ...
	}

	go app.server.Run()
}
