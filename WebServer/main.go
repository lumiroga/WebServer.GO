package main

func main() {
	server := NewServer(":3001")
	server.Handle("/", HandleRoot)
	server.Handle("/profile", server.AddMiddleware(HandleProfile, CheckAuth(), Login()))
	server.Listen()
}
