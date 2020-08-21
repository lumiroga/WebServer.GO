package main

func main() {
	server := NewServer(":3001")
	server.Handle("/", HandleRoot)
	server.Handle("/profile", HandleProfile)
	server.Listen()
}
