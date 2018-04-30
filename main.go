package main

import (
   "os"
   "net/http"
   "io"
)

func main() {
   http.HandleFunc("/", handler)
   http.ListenAndServe(":8999",nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
   name, err := os.Hostname()
   if err != nil {
      panic(err)
   } 
   io.WriteString(w, "Build 18 : Hello world! from node " + name )
}
