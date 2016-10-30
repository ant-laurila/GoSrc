package main

import (
	"fmt"
	"math"
	"strings"
)


func addTwo(x int, y int) int {
	return x + y
}

func ord(s string) string {
	a := strings.Split(s, "\t")
	q := []string{a[2], a[1], a[0]}
	
	return strings.Join(q, ":")
}

func makeCit(s string) string {
  a := strings.Split(s, "\t")
  q := "Citationtest " + a[2] + ". Made by " + a[0] + "-" + a[1]
  return q
}

func main() {
	var nro [2]int
	nro[0]=4
	nro[1]=6
	
	i := 12345
	
    fmt.Printf("hello, world\n")
	fmt.Println("Kultainto Piin tunnus on ", math.Pi)
	fmt.Printf("Tee yhteenlasku: %d + %d=", i, nro[1])
	fmt.Println("=>", addTwo(i,nro[1]))
	fmt.Println("######################################")
	fmt.Printf("Muuttuja i on tyyppiä %T", i)
    fmt.Println("######################################")
	
	
	o := "Jotain\tmuuta\ttäällä"
	fmt.Printf("%v", o)
	fmt.Println("######################################")
	fmt.Println(ord(o))
	fmt.Println("######################################")
    fmt.Println(makeCit(o))
	} 
