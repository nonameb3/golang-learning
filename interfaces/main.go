package main

import "fmt"

type bot interface {
	getGreeting() string
}

type botEnglish struct{}
type botThai struct{}

func main() {
	bEN := botEnglish{}
	bTH := botThai{}

	printGreeting(bEN)
	printGreeting(bTH)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (b botEnglish) getGreeting() string {
	return "Hello World"
}

func (b botThai) getGreeting() string {
	return "สวัสดี"
}
