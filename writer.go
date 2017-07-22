package main

import "os"
import "io"
import "fmt"
import "strings"

import "github.com/codecat/go-libs/log"

func addSlashes(str string) string {
	ret := str
	ret = strings.Replace(ret, "\\", "\\\\", -1)
	ret = strings.Replace(ret, "\"", "\\\"", -1)
	return ret
}

func writeField(f io.Writer, class ClassInfo, field FieldInfo) {
	fmt.Fprintf(f, "        { ")

	fmt.Fprintf(f, "ValueType::%s", translateTokenType(field.Type))
	fmt.Fprintf(f, ", \"%s\"", field.Name)
	if field.DefaultValue != "" {
		fmt.Fprintf(f, ", \"%s\"", addSlashes(field.DefaultValue))
	} else {
		fmt.Fprintf(f, ", nullptr")
	}
	fmt.Fprintf(f, ", offsetof(%s, %s)", class.Name, field.Name)

	fmt.Fprintf(f, " },\n")
}

func writeClass(f io.Writer, class ClassInfo) {
	fmt.Fprintf(f, "  class Tables_%s {\n", class.Name)
	fmt.Fprintf(f, "  public:\n")
	fmt.Fprintf(f, "    Tables_%s() {\n", class.Name)
	fmt.Fprintf(f, "      ClassInfo &ci = ncg::Classes.add();\n")
	fmt.Fprintf(f, "      ci.Name = \"%s\";\n", class.Name)
	fmt.Fprintf(f, "      ci.Fields = {\n")

	for _, field := range class.Fields {
		writeField(f, class, field)
	}

	fmt.Fprintf(f, "      };\n")
	fmt.Fprintf(f, "      ci.Methods = {\n")
	fmt.Fprintf(f, "        //TODO\n")
	fmt.Fprintf(f, "      };\n")
	fmt.Fprintf(f, "    }\n")
	fmt.Fprintf(f, "  };\n")
	fmt.Fprintf(f, "  static Tables_%s _Tables_%s;\n", class.Name, class.Name)
}

func writeResults(fnmSource, fnm string) bool {
	log.Info("Writing results to %s", fnm)

	f, err := os.Create(fnm)
	if err != nil {
		log.Error("Failed to create output file %s: %s", fnm, err.Error())
		return false
	}
	defer f.Close()

	fmt.Fprintf(f, "#include <ncg.h>\n")
	fmt.Fprintf(f, "#include \"%s\"\n", fnmSource)
	fmt.Fprintf(f, "namespace ncg { namespace generated {\n")

	for _, class := range results {
		writeClass(f, class)
	}

	fmt.Fprintln(f, "} }")
	return true
}
