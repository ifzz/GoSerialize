package code

import (
	"fmt"
	"strings"
)

type StructField struct {
	name string
	typeName string
}

func GenPackageHeaderAndImports(name string) string {
	return fmt.Sprintf(`
		package %s
		import (
			"serialize/encoders"
			"serialize/decoders"
			"errors"
			"fmt"
		)
	
		// THIS FILE WAS GENERATED BY serialize
		// PLEASE DO NOT EDIT
	`, name)
}

func GenSerializationHeader(name string) string {
	return fmt.Sprintf(`
		func (self %s) Serialize() ([]byte, error) {
			var output, bytesTemp []byte
	`, name)
}

func GenSerializationFooter() string {
	return "return output, nil }"
}

func GenFieldSerialization(fieldInfo StructField) string {
	return fmt.Sprintf(`
		bytesTemp = encoders.%sAsBytes(self.%s)
		output = append(output, bytesTemp...)
	`, strings.Title(fieldInfo.typeName), fieldInfo.name)
}

func GenUnserializationHeader(name string) string {
	return fmt.Sprintf(`
		func (self %s) Unserialize(data []byte) (interface{}, error) {
			var output %s
			var index uint64 = 0
			var consumed uint64 = 0
			var err error
	`, name, name)
}

func GenFieldUnserialization(fieldInfo StructField) string {
	return fmt.Sprintf(`
		output.%s, consumed, err = decoders.%sFromBytes(data[index:])
		if err != nil {
			return output, errors.New(fmt.Sprintf("Could not decode at %%d: %%s\n", index, err.Error()))
		}
		index += consumed
	`, fieldInfo.name, strings.Title(fieldInfo.typeName))
}

func GenUnserializationFooter() string {
	return "return output, nil }"
}