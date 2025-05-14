package main 

import (
  "fmt"
  "log"
  "net/http"
)

func formHandler(w http.ResponseWriter, r * http.Request){
  if err:= r.ParseForm(); err != nil {
    fmt.Fprintf(w, "Parseform() err: %v", err)
    return
  }
  fmt.Fprintf(w, "POST request is successfull!")
  name := r.FormValue("name")
  address:= r.FormValue("address")
  fmt.Fprintf(w, "The name is %s , lives in %s", name, address)
}


func helloHandler(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/hello"{
    http.Error(w, "404 Page not found", http.StatusNotFound)
    return
  }
  if r.Method != "GET"{
    http.Error(w , "405 Method Not Allowed", http.StatusNotFound)
    return 
  }
  fmt.Fprintf(w, "/hello!")
}



func main(){
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  http.HandleFunc("/form", formHandler)
  http.HandleFunc("/hello", helloHandler)

  fmt.Printf("Starting server port at 8080")
  if err:= http.ListenAndServe(":8080", nil); err!=nil{
    log.Fatal(err)
  }
} 
