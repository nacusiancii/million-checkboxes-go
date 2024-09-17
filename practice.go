package main

import(
	"fmt"
	"net/http"
)
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

//This is a go comment
/*
 * This is a multiline comment
 */
func handleAreaOfRectangle(w http.ResponseWriter, r *http.Request){
	length := 8
	breadth := 4
	area := length*breadth
	/*
 	 *String formating:
   	 *%d->int
	 *%t->boolean
	 *%s->string
  	 *%f->float
	 */
 	fmt.Fprintf(w, "Area of Rectangle: %d\n", area)
	var isLargeRectangle bool = (area>=200)
	fmt.Fprintf(w, "Large Rectangle: %t\n",isLargeRectangle)
}
