// abbreviation-checker checks the names in a file of code for possible abbreviations.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/BluntSporks/abbreviation"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Parsing flags.
	lang := flag.String("lang", "", "Name of language whose keywords should be ignored")
	file := flag.String("file", "", "Name of source code file to check")
	short := flag.String("short", "", "Short word to check if valid abbreviation of long")
	long := flag.String("long", "", "Long word to check if short is valid abbreviation of it")
	flag.Parse()

	// Check arguments.
	if *file != "" {
		checkFile(*file, *lang)
	} else if *short != "" && *long != "" {
		checkWord(strings.ToLower(*short), strings.ToLower(*long))
	} else {
		log.Fatal("Bad arguments")
	}
}

// checkFile checks a file.
func checkFile(file string, lang string) {
	// Match words.
	wordRegExp := regexp.MustCompile(`\w+`)

	// Ignore strings and line comments.
	ignoreRegExp := regexp.MustCompile(`".*?"|//.*`)

	// Open file.
	hdl, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()

	// Scan file line by line, removing dupes.
	scanner := bufio.NewScanner(hdl)

	// Keep track of whether the last line was empty or not.
	filePrinted := false
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		bare := ignoreRegExp.ReplaceAllString(strings.ToLower(line), "")
		if !filePrinted {
			// Add filename so it gets checked too.
			bare = strings.ToLower(file) + " " + bare
		}
		words := wordRegExp.FindAllString(bare, -1)
		matches := make(map[string]string)
		for _, word := range words {
			short := abbr.LookUp(word)
			if short != "" && !abbr.IsKeyword(lang, word) {
				matches[word] = short
			}
		}
		if len(matches) > 0 {
			if !filePrinted {
				fmt.Printf("File: %s\n", file)
				filePrinted = true
			}
			fmt.Printf("Line: %s\n", strings.TrimSpace(line))
			for long, short := range matches {
				fmt.Printf("%s => %s\n", long, short)
			}
			fmt.Println()
		}
	}
}

// checkWord checks a word.
func checkWord(short string, long string) {
	pass := true
	if !abbr.IsLongEnough(short) {
		fmt.Println("Fail: Short form should be at least 3 characters long")
		pass = false
	}
	if abbr.IsAcronym(short, long) {
		if pass {
			fmt.Println("Pass, because good acronym")
		}
	} else {
		if !abbr.AllLetters(short, long) {
			fmt.Println("Fail: All letters of long form not found in same order in short form")
			pass = false
		}
		if !abbr.HasVowel(short) {
			fmt.Println("Fail: Short form should have at least one vowel or y to make it pronounceable")
			pass = false
		}
		if !abbr.IsShortEnough(short, long) {
			fmt.Println("Fail: Short form should be no more than half as long as long form")
			pass = false
		}
		if pass {
			fmt.Println("Pass, because good abbreviation")
		}
	}
}
