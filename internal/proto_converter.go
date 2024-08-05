package internal

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/catalystcommunity/app-utils-go/errorutils"
)

type ProtoConverter struct {
	Object      string
	RawFieldMap map[string]string
}

func (c *ProtoConverter) SetObject(object string) {
	c.Object = object
}

func (c *ProtoConverter) SetRawFieldMap(rawFieldMap map[string]string) {
	c.RawFieldMap = rawFieldMap
}

func (c *ProtoConverter) Convert() {
	convertedFieldMap := map[string]string{}
	convertedFieldSortedKeys := []string{}
	for fieldName, fieldType := range c.RawFieldMap {
		convertedFieldMap[fieldName] = SfProtoTypeMap[fieldType]
		convertedFieldSortedKeys = append(convertedFieldSortedKeys, fieldName)
	}
	sort.Strings(convertedFieldSortedKeys)
	// build file
	// write header first
	headerBuilder := strings.Builder{}
	// include graphql proto so that the graphql field names get specified
	headerBuilder.WriteString("syntax = \"proto3\";\n\nimport \"danielvladco/protobuf/graphql.proto\";\n\n")
	// write message
	messageBuilder := strings.Builder{}
	messageBuilder.WriteString(fmt.Sprintf("message %s {\n", c.Object))
	fieldNumber := 1
	for _, fieldName := range convertedFieldSortedKeys {
		// if it's the any type, import the google any type
		fieldType := convertedFieldMap[fieldName]
		if fieldType == ProtoAny || fieldType == ProtoBase64 {
			headerBuilder.WriteString(`import "google/protobuf/any.proto";\n`)
		}
		messageBuilder.WriteString(fmt.Sprintf("  %s %s = %d;\n", fieldType, fieldName, fieldNumber))
		fieldNumber++
	}
	messageBuilder.WriteString("}")
	// output to file
	file, err := os.Create(fmt.Sprintf("%s.proto", c.Object))
	defer file.Close()
	if err != nil {
		errorutils.LogOnErr(nil, "error writing to file", err)
		return
	}
	// write header
	_, err = file.WriteString(headerBuilder.String())
	if err != nil {
		errorutils.LogOnErr(nil, "error writing to file", err)
		return
	}
	// write message
	_, err = file.WriteString(messageBuilder.String())
	if err != nil {
		errorutils.LogOnErr(nil, "error writing to file", err)
		return
	}
}
