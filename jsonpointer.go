package gojsonschema

import (
	"errors"
	"strings"
)

var ErrInvalidContext = errors.New("invalid context")

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

//convertContextToJSONPointer convert context to json pointer
func convertContextToJSONPointer(context *JsonContext) []string {
	var result = []string{}
	buildJSONPointer(context, 0, &result)
	if len(result) == 0 || result[0] != STRING_ROOT_SCHEMA_PROPERTY {
		panic(ErrInvalidContext)
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

func (v *ResultErrorFields) setSubSchema(schema *subSchema) {
	v.schema = schema
}

func (v *ResultErrorFields) IsEmptySchema() bool {
	return v.schema == nil
}

func (v *ResultErrorFields) SchemaTitle() string {
	if v.schema == nil || v.schema.title == nil {
		return ""
	}
	return *v.schema.title
}
func (v *ResultErrorFields) SchemaDescription() string {
	if v.schema == nil || v.schema.description == nil {
		return ""
	}
	return *v.schema.description
}
func (v *ResultErrorFields) Pointer() []string {
	return convertContextToJSONPointer(v.Context())
}

func (v *ResultErrorFields) EncodedPointer() string {
	return EncodeJSONPointer(v.Pointer())
}
func walkSchema(sch *subSchema, depth int, data *[]string) {
	if sch.parent == nil {
		*data = make([]string, depth)
		return
	}
	depth++
	walkSchema(sch.parent, depth, data)
	(*data)[len(*data)-depth] = sch.property

}
func (v *ResultErrorFields) SchemaEncodedPointer() string {
	if v.schema == nil {
		return ""
	}
	var data []string
	walkSchema(v.schema, 0, &data)
	return EncodeJSONPointer(data)
}

// type Found struct {
// 	schema  *subSchema
// 	root    *Schema
// 	pointer []string
// }

// func (f *Found) Title() string {
// 	if f.schema == nil || f.schema.title == nil {
// 		return ""
// 	}
// 	return *f.schema.title
// }
// func (f *Found) Description() string {
// 	if f.schema == nil || f.schema.description == nil {
// 		return ""
// 	}
// 	return *f.schema.description
// }
// func (r *Found) Pointer() []string {
// 	return r.pointer
// }

// func (r *Found) EncodePointer() string {
// 	return EncodeJSONPointer(r.pointer)
// }
// func (r *Found) IsEmpty() bool {
// 	return r.schema == nil
// }

// func walkSchema(sch *subSchema, pointer []string) *subSchema {
// 	if len(pointer) == 0 {
// 		return sch
// 	}
// 	if !sch.types.IsTyped() {
// 		return nil
// 	} else {
// 		if sch.types.Contains(TYPE_OBJECT) {
// 			for _, v := range sch.propertiesChildren {
// 				if v.property == pointer[0] {
// 					return walkSchema(v, pointer[1:])
// 				}
// 			}
// 			return nil
// 		}
// 		if sch.types.Contains(TYPE_ARRAY) {
// 			_, err := strconv.Atoi(pointer[0])
// 			if err != nil {
// 				return nil
// 			}
// 			if len(sch.itemsChildren) > 0 {
// 				return (walkSchema(sch.itemsChildren[0], pointer[1:]))
// 			}
// 		}
// 	}
// 	return nil
// }

// func Find(schema *Schema, context *JsonContext) *Found {
// 	pointer := convertContextToJSONPointer(context)
// 	sch := walkSchema(schema.rootSchema, pointer)
// 	return &Found{
// 		schema:  sch,
// 		root:    schema,
// 		pointer: pointer,
// 	}
// }
