package main

import (
	"fmt"
)

func main() {
	numCh := make(chan string)
	alphaCh := make(chan string)
	done := make(chan bool)
	var terms int

	fmt.Println("Enter the no. of terms:")
	fmt.Scan(&terms)

	go Generator(terms, numCh, alphaCh)
	go Splitter(numCh, alphaCh, done)
	<-done
}

func Generator(terms int, NumOut chan<- string, AlphaOut chan<- string) {
	defer close(NumOut)
	defer close(AlphaOut)

	n := make([]int, terms)
	a := make([]byte, terms)

	for i := 1; i <= terms; i++ {
		n[i-1] = i
		alpha := 'A' + i - 1
		a[i-1] = byte(alpha)
	}

	for i := 0; i < terms; i = i + 2 {
		NumOut <- fmt.Sprintf("%d%d", n[i], n[i+1])
		AlphaOut <- string(a[i]) + string(a[i+1])
	}
}

func Splitter(numberCh <-chan string, alphaCh <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for {
		select {
		case val, ok := <-numberCh:
			if !ok {
				return
			}
			fmt.Printf("%v", val)
		case val, ok := <-alphaCh:
			if !ok {
				return
			}
			fmt.Printf("%v", val)
		}
	}
}
