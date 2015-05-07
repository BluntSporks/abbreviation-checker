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

Usage:

    abbreviation-checker [-lang=(go|php)] -file=FILENAME
    abbreviation-checker -short=SHORTFORM -long=LONGFORM

Options:

    -lang=(go|php)    Programming languages whose keywords should be ignored
    -file=FILENAME    Name of file to check
    -short=SHORTFORM  Short form of word to check
    -long=LONGFORM    Long form of word or acronym expansion to check against short form
