package main

import (
	"fmt"
	"strings"
)

type todoList struct {
	data []todo
}

type todo struct {
	id   int
	text string
}

func (t *todoList) add(id int, text string) {
	t.data = append(t.data, todo{id: id, text: text})
}

// type 1: search by single word
func (t *todoList) get2(keyword string) (output []int) {
	k := strings.Replace(keyword, " ", "+", -1)
	for _, item := range t.data {
		isFound := strings.Contains(item.text, k)
		if isFound {
			output = append(output, item.id)
		}
	}

	return
}

// type 2: search by single word
func (t *todoList) get(keyword string) (output []int) {
	k := strings.Split(keyword, " ")
	if len(k) > 1 {
		return
	}

	for _, item := range t.data {
		isFound := strings.Contains(item.text, k[0])
		if isFound {
			output = append(output, item.id)
		}
	}

	return
}

func (t *todoList) remove(id int) {
	for idx, item := range t.data {
		if id == item.id {
			t.data = append(t.data[:idx], t.data[idx+1:]...)
			break
		}
	}
}

func main() {
	tl := todoList{data: []todo{}}
	// ADD
	tl.add(1, " lorem ipsum dolor sit amet")
	tl.add(2, " consectetur adipiscing elit sed do eiusmod tempor")
	tl.add(3, " ipsum eiusmod lorem tempor")

	fmt.Println("========================== ADD ============================")

	for _, v := range tl.data {
		fmt.Printf(`add(%v, "%v")`, v.id, v.text)
		fmt.Println("")
	}

	// GET-Able to search by single word
	result1 := tl.get("lorem")
	result2 := tl.get("consectetur")
	result3 := tl.get("aliqua")
	result4 := tl.get("lorem ipsum")

	fmt.Printf(`
	======== GET-Able to search by single word =========
	get("lorem") return %v
	get("consectetur") return %v
	get("aliqua") return %v
	get("lorem ipsum") return %v
	`, result1, result2, result3, result4)
	fmt.Println("")

	// Remove
	tl.remove(3)
	result5 := tl.get("lorem")

	fmt.Printf(`
	=========== Remove ===========
	remove(3) 
	get("lorem") return %v
	`, result5)
	fmt.Println("")
}
