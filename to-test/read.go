package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type UserInfo struct {
	fname string
	lname string
}

func getFileName() (fName string, err error) {
	stdin := bufio.NewReader(os.Stdin)

	fName, err = os.Getwd()
	fmt.Printf("Enter file name (current dir - %s):", fName)
	fName, err = stdin.ReadString('\n')
	if err == nil {
		fName = strings.TrimSpace(fName)
	}
	return
}

func truncName(nm string, sz int) (nmT string) {
	if len(nm) > sz {
		nmT = nm[0:sz]
	} else {
		nmT = nm
	}
	return nmT
}

func getUsersFromFile(fl *os.File, users *([]UserInfo), errStr *([]string)) (err error) {
	var uI UserInfo
	var str string

	stdin := bufio.NewReader(fl)
	for err != io.EOF {
		str, err = stdin.ReadString('\n')
		if err == nil || (err == io.EOF && len(str) > 0) {
			str = strings.TrimSpace(str)
			strS := strings.Split(str, " ")
			if len(strS) == 2 {
				uI.fname = truncName(strS[0], 20)
				uI.lname = truncName(strS[1], 20)
				*users = append(*users, uI)
			} else {
				*errStr = append(*errStr, str)
			}
		}
	}

	if err == io.EOF {
		err = nil
	}
	return err
}

func main() {
	var fName string
	var err error
	var users []UserInfo
	var errStr []string
	var fl *os.File
	var uI UserInfo
	var valS string

	fName, err = getFileName()
	if err == nil {
		fl, err = os.OpenFile(fName, os.O_RDONLY, 0x333)
		if err == nil {
			defer fl.Close()
			fmt.Printf("Processing file %s...\n", fName)
			err = getUsersFromFile(fl, &users, &errStr)
		} else {
			fmt.Println("Error open file: ", err)
		}
	}

	if err == nil { // to close the file faster
		fmt.Println("Recognized users:")
		for _, uI = range users {
			fmt.Printf(" %s %s\n", uI.fname, uI.lname)
		}
		fmt.Println("\nUnrecognized strings:")
		for _, valS = range errStr {
			fmt.Printf(" '%s'\n", valS)
		}
	}
}
