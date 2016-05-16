package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    location:=fmt.Sprintf("https://%s%s",r.Host,r.URL);
    http.Redirect(w,r,location,302);
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,"everything is fine.")
}

func main() {
  mux1 := http.NewServeMux();
  mux1.HandleFunc("/",handler);
  redirectionServer := &http.Server{
    Addr: ":5500",
    Handler: mux1,
  }
  go redirectionServer.ListenAndServe();

  mux2 := http.NewServeMux();
  mux2.HandleFunc("/",healthHandler);
  healthcheckServer := &http.Server{
    Addr: ":8080",
    Handler: mux2,
  }
  fmt.Printf("\r\nserver is running.\r\n\r\n");
  healthcheckServer.ListenAndServe();
}
