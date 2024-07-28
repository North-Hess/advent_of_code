package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type node struct {
	data     *string
	nextNode *node
}

type linkedList struct {
	head *node
	tail *node
}

func (l *linkedList) add(value string) {
	newNode := &node{
		data:     &value,
		nextNode: nil,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.nextNode = newNode
		l.tail = newNode
	}
}

func main() {
	file, err := os.Open("2024/Golang/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineValues := linkedList{}
		for _, character := range scanner.Text() {
			if unicode.IsNumber(character) {
				lineValues.add(string(character))
			}
		}
		stringValue := *lineValues.head.data + *lineValues.tail.data
		convertedValue, err := strconv.Atoi(stringValue)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + convertedValue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
