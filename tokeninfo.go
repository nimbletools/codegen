package main

func isTokenType(token string) bool {
	switch token {
	case "int", "long", "float", "double", "bool": fallthrough
	case "int8_t", "int16_t", "int32_t", "int64_t": fallthrough
	case "uint8_t", "uint16_t", "uint32_t", "uint64_t": fallthrough

	case "void":
		return true
	}
	return false
}
