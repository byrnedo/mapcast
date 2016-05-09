package mapcast

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strconv"
	"strings"
)

type fieldNamer func(reflect.StructField) string

type MapCaster struct {
	timeFormat  string
	inputNamer  fieldNamer
	outputNamer fieldNamer
}

func NewMapCaster() *MapCaster {
	return new(MapCaster)
}

func (m *MapCaster) TimeFormat(format string) {
	m.timeFormat = format
}

func (m *MapCaster) StdInput() {
	m.inputNamer = stdFieldNamer
}

func (m *MapCaster) StdOutput() {
	m.outputNamer = stdFieldNamer
}

func (m *MapCaster) JsonInput() {
	m.inputNamer = jsonFieldNamer
}

func (m *MapCaster) BsonOutput() {
	m.outputNamer = bsonFieldNamer
}

func (m *MapCaster) ProtoInput() {
	m.inputNamer = protoFieldNamer
}

func (m *MapCaster) Cast(inMap map[string]string, target interface{}) map[string]interface{} {
	return cast(inMap, target, m.inputNamer, m.outputNamer)
}

func stdFieldNamer(field reflect.StructField) string {
	return field.Name
}

func jsonFieldNamer(field reflect.StructField) string {
	t := field.Tag.Get("json")
	tArr := strings.Split(t, ",")

	fieldName := field.Name

	if len(tArr) > 0 && len(tArr[0]) > 0 {
		switch tArr[0] {
		case "-":
			return ""
		default:
			return tArr[0]
		}
	}
	return strings.ToLower(fieldName)
}

func bsonFieldNamer(field reflect.StructField) string {
	t := field.Tag.Get("bson")
	tArr := strings.Split(t, ",")

	fieldName := field.Name

	if len(tArr) > 0 && len(tArr[0]) > 0 {
		return tArr[0]
	}
	return strings.ToLower(fieldName)
}

func protoFieldNamer(field reflect.StructField) string {
	t := field.Tag.Get("protobuf")
	tArr := strings.Split(t, ",")

	fieldName := field.Name

	for _, sec := range tArr {
		sec = strings.TrimSpace(sec)
		if strings.HasPrefix(sec, "name=") && len(sec) > 5 {
			return sec[5:]
		}

	}
	return fieldName
}

func Cast(inMap map[string]string, target interface{}) (outMap map[string]interface{}) {
	return cast(inMap, target, stdFieldNamer, stdFieldNamer)
}

func CastViaJson(inMap map[string]string, target interface{}) (outMap map[string]interface{}) {
	return cast(inMap, target, jsonFieldNamer, stdFieldNamer)
}

func CastViaJsonToBson(inMap map[string]string, target interface{}) (outMap map[string]interface{}) {
	return cast(inMap, target, jsonFieldNamer, bsonFieldNamer)
}

func CastViaProtoToBson(inMap map[string]string, target interface{}) (outMap map[string]interface{}) {
	return cast(inMap, target, protoFieldNamer, bsonFieldNamer)
}

func cast(inMap map[string]string, target interface{}, fieldNamer fieldNamer, fieldRenamer fieldNamer) (outMap map[string]interface{}) {
	outMap = make(map[string]interface{})

	structElems := reflect.TypeOf(target).Elem()
	structValues := reflect.ValueOf(target).Elem()

	for i := 0; i < structElems.NumField(); i++ {

		fieldName := fieldNamer(structElems.Field(i))

		if origVal, found := inMap[fieldName]; found == true {
			if iVal, err := stringToType(fmt.Sprintf("%s", origVal), structValues.Field(i).Interface()); err == nil {
				fieldName = fieldRenamer(structElems.Field(i))
				outMap[fieldName] = iVal
			}
		}
	}
	return
}

func stringToType(val string, valType interface{}) (interface{}, error) {
	switch valType.(type) {
	case bson.ObjectId:
		if bson.IsObjectIdHex(val) {
			return bson.ObjectIdHex(val), nil
		}
	case bool:
		return strconv.ParseBool(val)
	case string:
		return val, nil
	case int:
		return strconv.Atoi(val)
	case int8:
		return strconv.ParseInt(val, 10, 8)
	case int16:
		return strconv.ParseInt(val, 10, 16)
	case int32:
		return strconv.ParseInt(val, 10, 32)
	case int64:
		return strconv.ParseInt(val, 10, 64)
	case uint:
		newVal, err := strconv.Atoi(val)
		return uint(newVal), err
	case uint8:
		strconv.ParseUint(val, 10, 8)
	case uint16:
		strconv.ParseUint(val, 10, 16)
	case uint32:
		strconv.ParseUint(val, 10, 32)
	case uint64:
		strconv.ParseUint(val, 10, 64)
	case float32:
		iVal, err := strconv.ParseFloat(val, 32)
		return float32(iVal), err
	case float64:
		iVal, err := strconv.ParseFloat(val, 64)
		return float64(iVal), err
	default:
		return nil, errors.New("Type not handled")
	}
	return nil, errors.New("Not handled")

}
