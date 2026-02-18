package main

import (
	"dsa/arrays"
	
	"os"
)

func main() {
	if len(os.Args)!=3{
		return
	}
	s1:=os.Args[1]
	s2:=os.Args[2]

	output:=arrays.Hidenp(s1,s2)

	os.Stdout.WriteString(output+"\n")
}


