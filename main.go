package main

import (
	"fmt"
	"sync"

	"golang.org/x/exp/slices"
)

type List struct {
	Data []string
	Wait *sync.WaitGroup
}

func (l *List) myFunc(s string, m *sync.Mutex, re *[]string) {
	defer l.Wait.Done()

	m.Lock()
	c := slices.Contains(*re, s)
	// m.Unlock()

	if !c {
		// m.Lock()
		*re = append(*re, s)
		// m.Unlock()

		r := l.someText(s)

		// m.Lock()
		l.Data = append(l.Data, r...)
		// m.Unlock()
	}

	m.Unlock()

}

func Contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func (l *List) someText(s string) []string {
	if s == "Hi" {
		return []string{"1", "2", "3"}
	}
	if s == "Adam" {
		return []string{"7", "8", "9"}
	}
	return []string{"4", "5", "6"}
}

func main() {

	for i := 0; i < 600; i++ {
		li := List{
			Data: []string{},
			Wait: &sync.WaitGroup{},
		}

		m := sync.Mutex{}
		re := []string{}

		myList := []string{
			"Hi",
			"Hello",
			"Hi",
			"Hi",
			"Adam",
			"Adam",
			"Adam",
		}

		for _, x := range myList {
			li.Wait.Add(1)
			go li.myFunc(x, &m, &re)
		}

		li.Wait.Wait()

		if len(re) > 3 {
			fmt.Println(li.Data)
			fmt.Println(re)
		}
	}
}
