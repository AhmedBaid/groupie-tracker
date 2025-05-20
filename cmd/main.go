package main

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie/routes"
	"groupie/tools"
)

func main() {
	var err error
	// parse all the html file from the template folder to variable Tp
	tools.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("Error parsing templates: ", err)
		return
	}
	// Register handlers
	routes.Route()
	//  run   the server
	fmt.Println("Server running at http://localhost:8081/")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Server is already runing")
		return
	}
}
