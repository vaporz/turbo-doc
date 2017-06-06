package main

import (
	"os"
	"io"
	"text/template"
	"github.com/vaporz/turbo-example/yourservice/gen/proto"
	"reflect"
	"fmt"
	"strings"
)

func main() {
	i := new(proto.YourServiceServer)
	t := reflect.TypeOf(i).Elem()
	numMethod := t.NumMethod()
	items := make([]string, 0)
	for i := 0; i < numMethod; i++ {
		method := t.Method(i)
		numIn := method.Type.NumIn()
		for j := 0; j < numIn; j++ {
			argType := method.Type.In(j)
			argStr := argType.String()
			if strings.HasSuffix(argStr, "Request") {
				arr := strings.Split(argStr, ".")
				name := arr[len(arr)-1:][0]
				items = findItem(items, name, argType)
			}
		}
	}
	var list string
	for _, s := range items {
		list += s + "\n"
	}
	writeFileWithTemplate(
		"/Users/xiaozhang/goworkspace/src/github.com/vaporz/turbo-example/yourservice/gen/grpcfields.yaml",
		fieldsYaml,
		fieldsYamlValues{List: list},
	)
}

func findItem(items []string, name string, structType reflect.Type) []string {
	numField := structType.Elem().NumField()
	item := "  - " + name + "["
	for i := 0; i < numField; i++ {
		fieldType := structType.Elem().Field(i)
		if fieldType.Type.Kind() == reflect.Ptr && fieldType.Type.Elem().Kind() == reflect.Struct {
			arr := strings.Split(fieldType.Type.String(), ".")
			typeName := arr[len(arr)-1:][0]
			argName := fieldType.Name
			item += fmt.Sprintf("%s %s,", typeName, argName)
			items = findItem(items, typeName, fieldType.Type)
		}
	}
	item += "]"
	return append(items, item)
}

func writeWithTemplate(wr io.Writer, text string, data interface{}) {
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(wr, data)
	if err != nil {
		panic(err)
	}
}

func writeFileWithTemplate(filePath, text string, data interface{}) {
	f, err := os.Create(filePath)
	if err != nil {
		panic("fail to create file:" + filePath)
	}
	writeWithTemplate(f, text, data)
}

type fieldsYamlValues struct {
	List string
}

var fieldsYaml string = `grpc-fieldmapping:
{{.List}}
`
