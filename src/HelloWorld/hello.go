package main

import (
    "net/http"
	"fmt"
)

func hello(w http.ResponseWriter,req *http.Request){
    fmt.Println("Say H");
}

func main() {
    http.HandleFunc("/",hello)  
  
    //服务器要监听的主机地址和端口号  
    err := http.ListenAndServe("127.0.0.1:8081", nil)  
  
    if err != nil {  
        fmt.Println("ListenAndServe error: ", err.Error())  
    }  
}
