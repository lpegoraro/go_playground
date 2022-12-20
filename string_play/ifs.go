package main

import (
	"errors"
	"fmt"
)

type Tree struct {
	name  string
	grade uint
	next  *Tree
}

func New(value string) (*Tree, error) {
	if len(value) > 0 {
		return &Tree{name: value}, nil
	} else {
		return nil, errors.New("cannot create empty value")
	}

}

func main() {
	// Pessoas na fila do banco
	first := Tree{
		name:  "John",
		grade: 7,
	}
	second := Tree{
		name:  "Joan",
		grade: 8,
	}
	third := Tree{
		name:  "Paula",
		grade: 4,
	}
	fourth := Tree{
		name:  "Bárbara",
		grade: 10,
	}
	first.next = &second
	second.next = &third
	third.next = &fourth
	node := &first
	for {
		if node == nil {
			break
		}
		fmt.Printf("Aluno: %s, nota: %d\n", node.name, node.grade)
		node = node.next
	}
}
