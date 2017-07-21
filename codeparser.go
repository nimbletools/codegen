package main

import "text/scanner"

import "github.com/codecat/go-libs/log"

var scan *scanner.Scanner

func nextToken() (token string, pos string, eof bool) {
	tok := scan.Scan()
	if tok == scanner.EOF {
		return "", scan.Pos().String(), true
	}
	return scan.TokenText(), scan.Pos().String(), false
}

func peekRune() rune {
	return scan.Peek()
}

func showError(s, pos string) {
	log.Error("%s at %s", s, pos)
}

func showErrorEof(pos string) {
	log.Error("Unexpected eof at %s", pos)
}

func showErrorUnexpectedToken(t, pos, expected string) {
	log.Error("Unexpected token '%s' at %s (expected: %s)", t, pos, expected)
}
