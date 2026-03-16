//package the file belongs to

package main

import (
	"fmt"
	"net/http"
)

//function runs , when request hits the server
//w http.ResponseWriter -> writer used to send data back
// to  client (send response )
//
//fmt.Fprintln(w, "Hello") means send "Hello" to the browser
//
/*
r *http.Request

What it is:
Contains information about the request.

Example data inside r:

URL

headers

method (GET / POST)

body

query params

*

*/
func handler(w http.ResponseWriter, r * http.Request) {

	fmt.Fprintf(w, "hello  this is my go server!")

}

func main() {
	//this is the new  route registration //this is the reverse order now 
	
	//route registration
	http.HandleFunc("/", handler)
	//means URL path "/"  → run handler()

	// / routing connects url to function

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
	/*
				Why ports exist

		A computer can run multiple servers.

		Example:

		8080 → backend server
		3000 → frontend
		5432 → database

		Ports separate services.
	*/

}
