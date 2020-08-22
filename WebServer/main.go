package main

func main() {
	server := NewServer(":3001")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", GenericPostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("GET", "/profile", server.AddMiddleware(HandleProfile, CheckAuth(), Login()))
	server.Listen()
}
