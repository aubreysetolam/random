package main

import "fmt"

func main() {
  var ind int 
  fmt.Scan(&ind)
  fmt.Print(doFib(ind))
}

func doFib(index int) int{
  var left, right int = 0,1
  for x := 1; x < index; x++{
    right, left = right + left, right
  }
  return right
}

//func recFib(index int)
