package main

import (
  "github.com/go-martini/martini"
  "net"
)

func main() {
 
  ips := make([]string, 1)

  m := martini.Classic()
  m.Get("/", func(w http.ResponseWriter, r *http.Request) string {
    ip, _, _ := net.SplitHostPort(r.RemoteAddr)
    ips = append(ips, ip)
    return string(ips)
  })
  m.Run()

}