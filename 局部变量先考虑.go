package main

import "fmt"

/* 声明全局变量 */
var a int = 20

func main9() {
   /* 声明局部变量 */
   var a int = 10

   fmt.Printf ("结果： g = %d\n",  a)
}