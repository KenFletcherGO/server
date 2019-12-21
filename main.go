package main

import (
	"expvar"
	_ "expvar"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

var myCount = expvar.NewInt("my.count")
var myStatus = expvar.NewString("my.status")

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hellohtml", hellohtml)
	http.HandleFunc("/formsubmit", formsubmit)
	http.HandleFunc("/template", hellotemphandler)
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

const hellotemplate = `
<!DOCTYPE html>
<html>
	<head>
		<title>Template page</title>
	</head>
	<body>
		<h1>Hello, {{.Name}}!</h1>
	</body>
</html>
`

var hellotmpl = template.Must(template.New(".").Parse(hellotemplate))

func hellotemphandler(response http.ResponseWriter, request *http.Request) {
	myCount.Add(1)
	myStatus.Set("Good")
	hellotmpl.Execute(response, map[string]interface{}{
		"Name": "Bob",
	})
}
