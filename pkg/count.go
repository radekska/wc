package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Counted struct {
	File  string
	Count int
	err   error
}

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
			return errors.New(fmt.Sprintf("error reading file %s", err)), 0
		}
		count += countFunc(line)
	}
	return nil, count
}

func countMany(files []string, countFunc func(string) int) (error, []Counted) {
	var wg sync.WaitGroup

	countedFiles := make([]Counted, 0, len(files))
	ch := make(chan Counted, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			err, counted := count(f, countFunc)
			ch <- Counted{f, counted, err}
		}(file)
	}
	wg.Wait()
	close(ch)

	for c := range ch {
		countedFiles = append(countedFiles, c)
		if c.err != nil {
			return c.err, nil
		}
	}
	return nil, countedFiles
}

func CountLines(files []string) (error, []Counted) {
	return countMany(files, func(s string) int { return 1 })
}

func CountWords(files []string) (error, []Counted) {
	return countMany(files, func(s string) int { return len(strings.Fields(s)) })
}

func CountCharacters(files []string) (error, []Counted) {
	return countMany(files, func(s string) int { return len(s) })
}
