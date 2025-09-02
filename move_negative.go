package main

import "fmt"


func move_negatives(array [] int)[]int {
    result:=make([]int, 0,len(array))
    for _,v:=range array{
        if v>=0{
            result=append(result, v)
        }
    }
    for _, v:=range array {
        if v<0{
            result=append(result, v)
        }
    }
  return result 
}

func main() {
    my_arr:=[]int{-1,-2,-3,1,2,3,4,5 }
    fmt.Println(move_negatives(my_arr))
}
