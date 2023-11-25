package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

func countLines(file string) (error, int) {
	return count(file, func(s string) int { return 1 })
}

func countWords(file string) (error, int) {
	return count(file, func(s string) int { return len(strings.Fields(s)) })
}

func countCharacters(file string) (error, int) {
	return count(file, func(s string) int { return len(s) })
}

func main() {
	pflag.BoolP("lines", "l", false, "count lines")
	pflag.BoolP("words", "w", false, "count words")
	pflag.BoolP("characters", "c", false, "count characters")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	file := pflag.Arg(0)
	if file == "" {
		fmt.Println("Please provide a file")
		return
	}
	l, w, c := viper.GetBool("lines"), viper.GetBool("words"), viper.GetBool("characters")
	lines, words, characters := 0, 0, 0
	err := error(nil)
	if l {
		err, lines = countLines(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(lines)
	}
	if w {
		err, words = countWords(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(words)
	}
	if c {
		err, characters = countCharacters(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(characters)
	}
}
