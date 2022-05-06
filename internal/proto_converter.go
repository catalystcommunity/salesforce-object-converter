package internal

import (
	"fmt"
	"github.com/catalystsquad/app-utils-go/errorutils"
	"os"
	"strings"
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
	for fieldName, fieldType := range c.RawFieldMap {
		convertedFieldMap[fieldName] = SfProtoTypeMap[fieldType]
	}
	// build file
	// write header first
	headerBuilder := strings.Builder{}
	headerBuilder.WriteString("syntax = \"proto3\";\n")
	// write message
	messageBuilder := strings.Builder{}
	messageBuilder.WriteString(fmt.Sprintf("message %s {\n", c.Object))
	fieldNumber := 1
	for fieldName, fieldType := range convertedFieldMap {
		// if it's the any type, import the google any type
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
