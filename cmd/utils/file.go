package utils

import (
	"fmt"
	"os"
)

func Save(data string, filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fprintln(f, data)

	if err != nil {
		fmt.Println(err)
	}

	f.Close()
}

func CreateFile(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
}
