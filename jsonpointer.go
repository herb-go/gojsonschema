package gojsonschema

import (
	"strings"
)

func encodeReferenceToken(token string) string {
	step1 := strings.Replace(token, `~`, `~0`, -1)
	step2 := strings.Replace(step1, `/`, `~1`, -1)
	return step2
}

func buildJSONPointer(context *JsonContext, depth int, data *[]string) {
	if context == nil {
		*data = make([]string, depth)
		return
	}
	depth++
	buildJSONPointer(context.tail, depth, data)
	(*data)[len(*data)-depth] = context.head
}

//ConvertContextToJSONPointer convert context to json pointer
func ConvertContextToJSONPointer(context *JsonContext) []string {
	var result = []string{}
	buildJSONPointer(context, 0, &result)
	if len(result) == 0 || result[0] != STRING_ROOT_SCHEMA_PROPERTY {

	}
	return result[1:]
}

func EncodeJSONPointer(p []string) string {
	var result = ""
	for _, v := range p {
		result = result + "/" + encodeReferenceToken(v)
	}
	return result
}

type PointedSchema struct {
	schema *subSchema
}

// func PointSchema(schema *Schema, context *JsonContext) *PointedSchema {

// }
