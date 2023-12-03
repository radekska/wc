package pkg

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Counted struct {
	File  string `json:"file"`
	Count int    `json:"count"`
}

func orchestrate(files []string, output string, countFunc func(string) int) error {
	errCh := make(chan error)
	ch := make(chan Counted, len(files))

	go save(output, len(files), ch, errCh)
	go countMany(files, ch, errCh, countFunc)

	for {
		select {
		case err := <-errCh:
			return err
		}
	}
}

func save(output string, numFiles int, ch <-chan Counted, errCh chan<- error) {
	f, err := os.Create(output)
	if err != nil {
		errCh <- err
	}
	defer f.Close()

	countedFiles := make([]Counted, 0, numFiles)
	e := json.NewEncoder(f)
	for c := range ch {
		countedFiles = append(countedFiles, c)
	}
	errCh <- e.Encode(countedFiles)
}

func countMany(files []string, ch chan<- Counted, errCh chan<- error, countFunc func(string) int) {
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			err, counted := count(f, countFunc)
			if err != nil {
				errCh <- err
			}
			ch <- Counted{f, counted}
		}(file)
	}
	wg.Wait()
	close(ch)
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

func CountLines(files []string, output string) error {
	return orchestrate(files, output, func(s string) int { return 1 })
}

func CountWords(files []string, output string) error {
	return orchestrate(files, output, func(s string) int { return len(strings.Fields(s)) })
}

func CountCharacters(files []string, output string) error {
	return orchestrate(files, output, func(s string) int { return len(s) })
}
