package main

import "fmt"

func main(){
   var e string = english()
   var j string = japanese()
   var c string = chinese()
   
   fmt.Println(e)
   fmt.Println(j)
   fmt.Println(c)
}

func english() string {
    return "Hello, world"
}

func japanese() string {
    return "Hello, world"      /* <- ここがバグ */
    return "こんにちは、世界"
}

func chinese() string {
    return "你好,是世界"
}
