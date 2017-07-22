package main

import "os"
import "strings"
import "text/scanner"

import "github.com/codecat/go-libs/log"

var scanState *ScanState

var currentClass int = -1

func parseMetaClass() bool {
	// "meta class"

	token, pos, eof := nextToken()
	if eof {
		showErrorEof(pos)
		return false
	}

	currentClass = addClassInfo(token)

	scanState.PushScope(token)
	return true
}

func parseMetaField(typeName, name string) bool {
	// "meta int m_score"

	token, pos, eof := nextToken()
	if eof {
		showErrorEof(pos)
		return false
	}

	defVal := ""

	for token != ";" {
		token, pos, eof = nextToken()
		if token == ";" {
			break
		}
		defVal += token + " "
	}

	addFieldInfo(currentClass, FieldInfo{ typeName, name, strings.TrimSpace(defVal) })

	return true
}

func parseMetaMethod(returnTypeName, name string) bool {
	// "meta int GetNumber"

	addMethodInfo(currentClass, MethodInfo{ returnTypeName, name })

	return true
}

func parseMetaType(typeName string) bool {
	// "meta int"

	token, pos, eof := nextToken()
	if eof {
		showErrorEof(pos)
		return false
	}

	if peekRune() == '(' {
		return parseMetaMethod(typeName, token)
	} else {
		return parseMetaField(typeName, token)
	}

	return true
}

func parseMeta() bool {
	// "meta"

	token, pos, eof := nextToken()
	if eof {
		showErrorEof(pos)
		return false
	}

	if token == "class" {
		return parseMetaClass()
	}

	if !scanState.Scopes.Empty() {
		if token == "virtual" {
			//TODO: We can use this information
			token, pos, eof = nextToken()
			if eof {
				showErrorEof(pos)
				return false
			}
		} else if token == "inline" {
			//TODO: We can use this information
			token, pos, eof = nextToken()
			if eof {
				showErrorEof(pos)
				return false
			}
		}

		if isTokenType(token) {
			return parseMetaType(token)
		}
	}

	showErrorUnexpectedToken(token, pos, "class name, virtual keyword, typename")
	return false
}

func main() {
	if len(os.Args) == 1 {
		print("Usage: codegen <filename>")
		return
	}

	fnm := os.Args[1]
	fnmOutput := os.Args[1][:len(fnm)-2] + "_tables.cpp"

	f, err := os.Open(fnm)
	if err != nil {
		log.Error("Couldn't open file: %s", err.Error())
		return
	}

	log.Info("Parsing %s", fnm)

	scanState = &ScanState{}
	scanState.Init()

	scan = &scanner.Scanner{}
	scan.Filename = fnm
	scan.Init(f)

	for {
		token, pos, eof := nextToken()
		if eof {
			break;
		}

		if token == "meta" {
			parseMeta()
		} else if token == "{" {
			scanState.ScopeDepth++
		} else if token == "}" {
			scanState.ScopeDepth--
			if scanState.ScopeDepth < 0 || scanState.Scopes.Empty() {
				showError("Unexpected end of scope", pos)
			} else {
				scope := scanState.Scopes.Top().(ScanScope)
				if scope.Depth == scanState.ScopeDepth + 1 {
					scanState.PopScope()
				}
			}
		}
	}

	writeResults(fnm, fnmOutput)
}
