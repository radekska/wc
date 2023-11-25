package pkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func count(file string, countFunc func(string) int) (error, int) {
	f, err := os.Open(file)
	if err != nil {
		return err, 0
	}
	defer f.Close()

	r := bufio.NewReader(f)
	count := 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		count += countFunc(line)
	}
	return nil, count
}

func CountLines(file string) (error, int) {
	return count(file, func(s string) int { return 1 })
}

func CountWords(file string) (error, int) {
	return count(file, func(s string) int { return len(strings.Fields(s)) })
}

func CountCharacters(file string) (error, int) {
	return count(file, func(s string) int { return len(s) })
}
