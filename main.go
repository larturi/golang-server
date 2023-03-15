package main

func main() {
	server := NewServer(":3100")
	server.Listen()
}
