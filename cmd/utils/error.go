package utils

import (
	"fmt"
	"os"
	"time"
)

func HandleError(err error) {
	fmt.Println("An error has been encountered and logged")
	t := time.Now()
	name := fmt.Sprintf("./%d-%d-%d.log", t.Day(), t.Month(), t.Year())
	CreateFile(name)
	Save(fmt.Sprintf("Time: %s, Error: %s", t.Format(time.Kitchen), err.Error()), name)
	os.Exit(1)
}
