package main

import (
	"fmt"
	"math/rand"
	"strings"
	"bufio"
	"os"
	"time"
)


func addTwo(x int, y int) int {
	return x + y
}


func drawHelp(i int) {
	fmt.Printf("(%v)", i)
	for n := 0; n<i; n++ { fmt.Printf("o") }
	fmt.Printf("\n")
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
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().Unix())
	
	i := rand.Intn(11)
	j := rand.Intn(11)
	
    fmt.Printf("Tee yhteenlasku: %d + %d=\n", i, j)
	drawHelp(i)
	drawHelp(j)
	fmt.Print("Anna vastaus: ")
	inp1, _ := reader.ReadString('\n')	
	correct := addTwo(i,j)
	fmt.Printf("Sinä vastasit: %v ja oikea vastaus on %v\n\n", inp1, correct)

	fmt.Println("######################################")
	fmt.Println("######################################")
	fmt.Printf("Muuttuja i on tyyppiä %T\n", i)
    fmt.Println("######################################")
	
	
	o := "Isi\tteki\ttämän"
	fmt.Printf("%v\n", o)
	fmt.Println("######################################")
	fmt.Println(ord(o))
	fmt.Println("######################################")
    fmt.Println(makeCit(o))
	} 
