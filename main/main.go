package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func TailFile(fileName string) {

	tails, err := tail.TailFile(fileName, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Poll:      true,
		MustExist: false,
	})
	if err != nil {
		fmt.Println("open file is err", err)
		return
	}
	for {
		line, ok := <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("is line=:", line)
	}

}
func main() {

}
