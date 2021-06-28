package main

import (
	"log"

	ds "github.com/jason-adam/data-structures/pkg/ds"
)

func main() {
	pq := ds.NewMaxPQ(10)
	pq.Insert("V")
	pq.Insert("T")
	pq.Insert("R")
	pq.Insert("W")
	pq.Insert("C")
	pq.Insert("X")
	pq.Insert("Z")
	pq.Insert("A")
	pq.Insert("B")

	log.Println(pq.GetPQ())

	x := pq.DeleteMax()
	log.Println(x)

	log.Println(pq.GetPQ())

	y := pq.DeleteMax()
	log.Println(y)

	log.Println(pq.GetPQ())
}
