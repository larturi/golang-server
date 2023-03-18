package main

func main() {
	server := NewServer(":3100")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", HandlePostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(
		HandleHome,
		CheckAuth(),
		Loggin(),
	))
	server.Listen()
}
