package main

func main() {
	server := NewServer(":3100")
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleHome)
	server.Listen()
}
