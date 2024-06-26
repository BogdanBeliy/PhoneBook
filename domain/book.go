package domain

import "fmt"

type BookItem struct {
	Name    string
	Surname string
	Tel     string
}

var Data []BookItem

func Search(key string) *BookItem {
	for i, v := range Data {
		if v.Surname == key {
			return &Data[i]
		}
	}
	return nil
}

func List() {
	for _, v := range Data {
		fmt.Println(v)
	}
}
