package main

import(
  "fmt"
  "./hello"
)

func main(){
   var e string = hello.English()
   var j string = hello.Japanese()
   var c string = hello.Chinese()
   
   fmt.Println(e)
   fmt.Println(j)
   fmt.Println(c)
}

