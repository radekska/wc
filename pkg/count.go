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
	Err   error
}

type CountedFiles struct {
	Files []Counted
	mu    sync.Mutex
}

func NewCountedFiles(s int) *CountedFiles {
	return &CountedFiles{Files: make([]Counted, 0, s)}
}

func (cf *CountedFiles) Add(c Counted) {
	cf.mu.Lock()
	defer cf.mu.Unlock()
	cf.Files = append(cf.Files, c)
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

func countMany(files []string, countFunc func(string) int) *CountedFiles {
	var wg sync.WaitGroup

	countedFiles := NewCountedFiles(len(files))

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			err, counted := count(f, countFunc)
			countedFiles.Add(Counted{File: f, Count: counted, Err: err})
		}(file)
	}
	wg.Wait()

	return countedFiles
}

func CountLines(files []string) *CountedFiles {
	return countMany(files, func(s string) int { return 1 })
}

func CountWords(files []string) *CountedFiles {
	return countMany(files, func(s string) int { return len(strings.Fields(s)) })
}

func CountCharacters(files []string) *CountedFiles {
	return countMany(files, func(s string) int { return len(s) })
}
