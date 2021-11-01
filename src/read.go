package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	Fname string    `json:"fname"`
	Lname string    `json:"lname"`
}

func (n Name) String() string {
	return fmt.Sprintf("[First: %s, Last: %s]", n.Fname, n.Lname)
}

func main() {
	var fileName string

	fmt.Println("Enter a file to read...")
	num1, _ := fmt.Scan(&fileName)
	_ = num1

	dat, e := ioutil.ReadFile(fileName)

	if e != nil {
		fmt.Println("Had an error: ", e)
		return
	}

	contents := string(dat)
	lines := strings.Split(contents, "\r\n")

	var arr = make([]Name, 0, 30)
	for _, line := range lines {
		names := strings.Split(line, " ")
		arr = append(arr, Name{Fname: names[0], Lname: names[1]})
	}

	barr, _ := json.Marshal(arr)
	fmt.Println(string(barr))
}