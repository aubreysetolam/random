package main

import "fmt"

func main() {
  var ind int 
  fmt.Scan(&ind)
  fmt.Print(doFib(ind))
  fmt.Print(recursiveFib(ind))
}

func doFib(index int) int{
  var left, right int = 0,1
  for x := 1; x < index; x++{
    right, left = right + left, right
  }
  return right
}

func recursiveFib(index int) int{
  if index <= 2{
    return 1
  }else{
    return recursiveFib(index-1) + recursiveFib(index-2)
  }
}
