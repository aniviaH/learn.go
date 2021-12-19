package main

import (
	"fmt"
	"os"
	"time"
)

func deferGuess() {
	startTime := time.Now()
	defer fmt.Println("duration:", time.Now().Sub(startTime))
	defer func() {
		fmt.Println("duration upgraded:", time.Now().Sub(startTime))
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("startTime:", startTime)
}

func openFile() {
	//filename := "/test.txt"
	filename := "/"
	fileObj, err := os.Open(filename)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	defer fileObj.Close()

	defer func() {
		fileObj.Close()
	}()

	fileObj = nil
	fileObj.Close()
}
