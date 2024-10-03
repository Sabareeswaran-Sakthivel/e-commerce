package routes

/*
Init initializes the server by getting the configuration,
creating a new router, and running the server on the specified port.
*/
func Init() {
	r := NewRouter()

	port := "localhost:7075"

	r.Run(port)
}
