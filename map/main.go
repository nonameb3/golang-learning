package main

import "fmt"

func main() {
	data := map[string]string{
		"dog": "bark",
		"cat": "purr",
	}

	updateMap(data)
	printMap(data)
}

func updateMap(m map[string]string) {
	m["bird"] = "my bird"
	m["baby"] = "my baby"
	delete(m, "baby")
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, "name: ", value)
	}
}
