package main

import (
  "fmt"
  "net/http"
  "crypto/tls"
  "io/ioutil"
  "encoding/json"
  //"strings"
)

func main() {
  //url := "https://10.4.45.77/api/json/types/lun-maps/"
  url := "https://10.4.45.77/api/json/types/volumes/"
  user := "admin"
  pass := "password"

  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  req, err := http.NewRequest("GET", url, nil)
  req.SetBasicAuth(user, pass)
  cli := &http.Client{Transport: tr}
  rep, err := cli.Do(req)
  fmt.Println("INFO: rep is> ", rep)
  fmt.Println("INFO: err is> ", err)

  defer rep.Body.Close()
  body, err := ioutil.ReadAll(rep.Body)
  //fmt.Println("INFO; body is> ", body)

  type Payload struct {
    volumes, links string
  }

  var payload Payload
  json.Unmarshal(body, &payload) 
  fmt.Printf("%+v", payload)

}
