package main

import "fmt"


func reverse_array(array [] int){
    for i,j:=0,len(array)-1;i<j;i,j=i+1,j-1{
        array[i],array[j]=array[j], array[i]
    }
}
func main() {
    my_array:=[]int{1,2,3,4,5}
    reverse_array(my_array)
    fmt.Println(my_array)

}
