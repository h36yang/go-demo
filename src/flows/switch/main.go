package main

type httpRequest struct {
	Method string
}

func main() {
	r := httpRequest{Method: "GET"}

	// NOTE: Go has implicit break statements among the cases by default
	// If we want to case to fall through to the next, use explicit "fallthrough" keyword
	switch r.Method {
	case "GET":
		println("GET request")
	case "POST":
		println("POST request")
	case "PUT":
		println("PUT request")
	case "DELETE":
		println("DELETE request")
	default:
		println("Unhandled method")
	}
}
