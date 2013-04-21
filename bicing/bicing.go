package bicing

import (
  "net/http"
  "io/ioutil"
)

const URL string = "https://www.bicing.cat/localizaciones/getJsonObject.php"

func Stations() (body []byte) {
  response, _ := http.Get(URL)

  defer response.Body.Close()
  body, _ = ioutil.ReadAll(response.Body)

  return
}
