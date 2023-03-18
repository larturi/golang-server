package main

func main() {
	server := NewServer(":3100")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMiddleware(
		HandleHome,
		CheckAuth(),
		Loggin(),
	))
	server.Listen()
}
