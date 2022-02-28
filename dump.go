package dbg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Dump2String like Dump, but JSON string is returned
func Dump2String(v interface{}) string {

	sb := &strings.Builder{}

	if false {
		var strValue reflect.Value = reflect.ValueOf(&v) // take pointer
		indirectStr := reflect.Indirect(strValue)        // pointer back to value
		typeAsString := reflect.MakeMap(indirectStr.Type())
		fmt.Fprintf(sb, "\ttype:  %v\n", typeAsString)
	}
	fmt.Fprintf(sb, "\n\ttype  %T:\n", v)

	firstColLeftMostPrefix := " "
	bts, err := json.MarshalIndent(v, firstColLeftMostPrefix, "\t")
	if err != nil {
		s := fmt.Sprintf("error indent: %v\n", err)
		return s
	}

	bts = bytes.Replace(bts, []byte(`\u003c`), []byte("<"), -1)
	bts = bytes.Replace(bts, []byte(`\u003e`), []byte(">"), -1)
	bts = bytes.Replace(bts, []byte(`\n`), []byte("\n"), -1)

	fmt.Fprint(sb, string(bts))

	return sb.String()
}

// Dump converts to JSON string and prints to standard logger;
// sort of Dump2Log
func Dump(v interface{}) {
	log.Print(Dump2String(v))
}

// Dump2Pre like Dump, , but JSON string is returned nested in <pre>...</pre>
func Dump2Pre(v interface{}, styles ...string) string {

	style := ""
	if len(styles) > 0 {
		style = styles[0]
	}

	sb := &strings.Builder{}
	fmt.Fprintf(sb, "<pre style='%v'>\n", style)
	fmt.Fprint(sb, Dump2String(v))
	fmt.Fprint(sb, "</pre>\n")
	return sb.String()

}
