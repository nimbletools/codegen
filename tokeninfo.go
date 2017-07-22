package main

import "github.com/codecat/go-libs/log"

func translateTokenType(token string) string {
	switch token {
	case "int8_t", "char": return "Int8"
	case "int16_t": return "Int16"
	case "int32_t", "int": return "Int32"
	case "int64_t", "long": return "Int64"
	case "float": return "Float"
	case "double": return "Double"
	case "bool": return "Bool"

	case "uint8_t": return "Uint8"
	case "uint16_t": return "Uint16"
	case "uint32_t": return "Uint32"
	case "uint64_t": return "Uint64"

	case "void": return "Void"
	}

	log.Warn("Unknown token type %s, using Int32 instead", token)
	return "Int32"
}

func isTokenType(token string) bool {
	switch token {
	case "char", "int", "long", "float", "double", "bool": fallthrough
	case "int8_t", "int16_t", "int32_t", "int64_t": fallthrough
	case "uint8_t", "uint16_t", "uint32_t", "uint64_t": fallthrough

	case "void":
		return true
	}
	return false
}
