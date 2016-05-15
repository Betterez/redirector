package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    location:=fmt.Sprintf("https://%s%s",r.Host,r.URL);
    http.Redirect(w,r,location,302);
}

func main() {
  http.HandleFunc("/",handler);
  port := ":5500"
  fmt.Printf("server starting at port %s",port);
  http.ListenAndServe(port, nil)
}
