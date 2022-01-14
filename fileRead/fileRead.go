package fileRead

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FileRead(u string) {
	fmt.Println("Hello,FileRead " + u)

	fp, e := os.Open(u)
	check(e)
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	read := ""

	start := time.Now()
	for scanner.Scan() {
		read = scanner.Text()
	}
	end := time.Now()

	diff := end.Sub(start)

	fmt.Printf("last row is " + read + "\n")
	fmt.Printf("result: %d milliseconds.\n", diff.Milliseconds())
}
