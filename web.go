package main

import (
  "os"
  "fmt"
  "net/http"
)

func main() {

  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", fs)
  bind :=fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
  fmt.Printf("listening on %s...", bind)
  http.ListenAndServe(bind, nil)

}
