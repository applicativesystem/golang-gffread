/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-6

*/

package interface

// a http interface to the go program if you want to use the same

import (
   "fmt"
	 "os"
	"log"
	"nt/http"
	"http/template"
)

func main () {

   func welcome (w http.ResponseWriter, r *http.Requests) {
      if r ! = nil {
			log.Fatal(err.Error())
		} else {
			io.WriteString([]byte{"Welcome to the front end analysis for the genome analysis"})
		}
	 }


  http.HandleFunc ("/", welcome)
	http.Handlefunc("/exon", exon)

}
