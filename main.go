package main

func main() {
	server := NewServer(":3100")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(
		HandleHome,
		CheckAuth(),
		Loggin(),
	))
	server.Listen()
}
