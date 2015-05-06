# abbreviation-checker
Golang program to suggest possible abbreviations of identifiers in source code.

## Purpose
To provide a utility to list possible abbreviations of identifiers in source code.

## Status
Ready to use

## Installation
This program is written in Google Go language. Make sure that Go is installed and the GOPATH is set up as described in
[How to Write Go Code](https://golang.org/doc/code.html).

The install this program and its dependencies by running:

    go get github.com/BluntSporks/abbreviation-checker

## Usage
The program runs in one of two modes:
* File checking, where it checks all the names in a file and suggests abbreviations.
* Word checking, where you check if an abbreviation follows all the rules for its long form.

### File checking
    abbreviation-checker -lang [language] -file [filename]

Arguments:
* lang: Use either go or php as language value to ignore keywords from that language. Argument is optional.
* file: Provide filename of source code to check.

### Word checking
    abbreviation-checker -short [short] -long [long]

Arguments:
* short: Short, or abbreviated, form that you are checking.
* long: Original long form that it is being checked against.
