package main

import (
    "net/http" //oh 不能用,分割
    "fmt"
)
//当遇到问题的时候，go停止正常输出？
func main(){
    testgetquestbaidu()
    fmt.Println("hello world")
}

func testgetquestbaidu(){
    url = "http://www.baidu.com"
    resp, err := http.Get(url)
    if err != nil{
        fmt.Errorf(err)
    } else {
        fmt.Print(resp)
    }
}
