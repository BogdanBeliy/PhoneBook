package main

import (
	"PhoneBook/domain"
	"fmt"
	"log"
	"os"
	"path"
)

func init() {
	domain.Data = append(domain.Data, domain.BookItem{Name: "Bogdan", Surname: "Beliy", Tel: "1234"})
	domain.Data = append(domain.Data, domain.BookItem{Name: "Vasya", Surname: "Pupkin", Tel: "12345"})
	domain.Data = append(domain.Data, domain.BookItem{Name: "Petya", Surname: "Sidorov", Tel: "123456"})
}

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatal("Not enogph arguments")
	}
	exe := path.Base(args[0])
	fmt.Printf("Usage: %s search | list <args>\n", exe)
	switch args[1] {
	case "search":
		if len(args) != 3 {
			fmt.Printf("Usage: search Surname %s")
			return
		}
		result := domain.Search(args[2])
		if result == nil {
			fmt.Printf("Item not found %s", args[2])
		}
		fmt.Printf("Name: %s, Surname: %s, Tel: %s", result.Name, result.Surname, result.Tel)
	case "list":
		domain.List()
	default:
		fmt.Println("Not a valid option")
	}
	return
}
