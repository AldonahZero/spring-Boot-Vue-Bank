package main

import (
    "gopkg.in/mgo.v2"
    "fmt"
)
func main() {
    session, err := mgo.Dial("mongodb://127.0.0.1:27017")
    defer session.Close()
    if err !=nil {
        fmt.Println(err)
        return
    }
    names,err:=session.DatabaseNames();
    if err !=nil {
        fmt.Println("未查询到数据库名字:",err)
    }
    fmt.Println(names)
}