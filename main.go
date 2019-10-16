package main

import (
	"expvar"
	"fmt"
	"io"
	"net/http"
	"text/template"

	_ "expvar"
)

var myCount = expvar.NewInt("my.count")

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hellohtml", hellohtml)
	http.HandleFunc("/fromsubmit", formsubmit)
	http.HandleFunc("/template", hellotemphandler)
	http.ListenAndServe(":9000", nil)
}

func hello(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	output := []byte("Hello " + name)
	fmt.Println(name, "says hello")
	response.Write(output)
}

func hellohtml(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	// output := []byte("<html><body><h1>Hello There!</h1></body></html>")
	// response.Write(output)
	io.WriteString(response, `
	
	`)
}

func formsubmit(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Welcome", request.FormValue("user"))
	fmt.Println("Your password is", request.FormValue("password"))
}

// template so its always the same
const hellotemplate = `
<!DOCTYPE html>
<html>
	<head>
	<title>Template page</title>
	</head>
	<body>
		<h1> Hello, {{.Name}}!</h1>
	</body>
</html>
`

var hellotmpl = template.Must(template.New(".").Parse(hellotemplate))

func hellotemphandler(response http.ResponseWriter, request *http.Request) {
	myCount.Add(1)
	hellotmpl.Execute(response, map[string]interface{}{
		"Name": "Bob",
	})
}
