package main

import "github.com/codecat/go-libs/log"

type FieldInfo struct {
	Type string
	Name string
	DefaultValue string
}

type MethodInfo struct {
	ReturnType string
	Name string
}

type ClassInfo struct {
	Name string
	Fields []FieldInfo
	Methods []MethodInfo
}

var results []ClassInfo

func addFieldInfo(class int, fi FieldInfo) {
	if class < 0 || class >= len(results) {
		log.Error("Tried adding field info for unset class index %d", class)
		return
	}

	if fi.DefaultValue == "" {
		log.Trace("Adding field info: '%s %s::%s'", fi.Type, results[class].Name, fi.Name)
	} else {
		log.Trace("Adding field info: '%s %s::%s = %s'", fi.Type, results[class].Name, fi.Name, fi.DefaultValue)
	}

	results[class].Fields = append(results[class].Fields, fi)
}

func addMethodInfo(class int, mi MethodInfo) {
	if class < 0 || class >= len(results) {
		log.Error("Tried adding method info for unset class index %d", class)
		return
	}

	log.Trace("Adding method info: '%s %s::%s(...)'", mi.ReturnType, results[class].Name, mi.Name)

	results[class].Methods = append(results[class].Methods, mi)
}

func addClassInfo(name string) int {
	log.Trace("Adding class info: '%s'", name)

	ret := len(results)
	results = append(results, ClassInfo{ name, []FieldInfo{}, []MethodInfo{} })
	return ret
}
