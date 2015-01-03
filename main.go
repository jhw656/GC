package main

import "fmt"

func main() {
	p:= []int{2,4,6,8,10}
	for i:=0;i<len(p);i++ {
	fmt.Printf("p[%d] = %d\n",i,p[i])
	}
}