package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hellohtml", hellohtml)
	http.HandleFunc("/formsubmit", formsubmit)
	http.ListenAndServe(":9000", nil)
}

func hello(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	output := []byte("Hello " + name)
	fmt.Println(name, "says hello!")
	response.Write(output)
}

func hellohtml(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	//output := []byte("<html><body><h1>Hello There!</h1></body></html>")
	//response.Write(output)
	io.WriteString(response, `
	<DOCTYPE html>
	<html>
	<head>
		<title>My Page</title>
	</head>
	<body>
		<h2>Welcome to My Page</h2>
		<p>This is a test of a go server</p>
	</body>
	</html>
	`)
}

func formsubmit(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Welcome", request.FormValue("user"))
	fmt.Println("Your password is", request.FormValue("password"))
}
