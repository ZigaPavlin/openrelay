package main

import (
  "github.com/notegio/0xrelay/types"
  "io/ioutil"
  "encoding/json"
  // "encoding/hex"
  // "reflect"
)

func main() {
  order := types.Order{}
  if orderData, err := ioutil.ReadFile("formatted_transaction.json"); err == nil {
    if err := json.Unmarshal(orderData, &order); err != nil {
      println(err.Error())
      return
    }
  } else {
    println(err.Error())
    return
  }
  println(order.Signature.Verify(order.Maker))
}
