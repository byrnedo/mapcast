package mapcast

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type inputStruct struct {
	String   string        `json:"jstring" bson:"bstring" protobuf:"bytes,1,opt,name=pstring"`
	Int      int           `json:"jint" bson:"bint" protobuf:"bytes,1,opt,name=pint"`
	Uint     uint          `json:"juint" bson:"buint" protobuf:"bytes,1,opt,name=puint"`
	ObjectId bson.ObjectId `json:"jobjectid" bson:"bobjectid" protobuf:"bytes,1,opt,name=pobjectid"`
}

func TestStdMapCast(t *testing.T) {

	inputData := map[string]string{
		"String":   "string",
		"Int":      "-1",
		"Uint":     "2",
		"ObjectId": bson.NewObjectId().Hex(),
	}

	caster := NewMapCaster()
	caster.StdInput()
	caster.StdOutput()

	targetStruct := inputStruct{}
	outputMap := caster.Cast(inputData, &targetStruct)

	expectedOutput := map[string]interface{}{
		"String":   "string",
		"Int":      -1,
		"Uint":     uint(2),
		"ObjectId": bson.ObjectIdHex(inputData["ObjectId"]),
	}

	for key, val := range expectedOutput {
		if gotVal, found := outputMap[key]; found == true {
			if gotVal == val {
				t.Log("Value matches:", key, val, gotVal)
				continue
			}
			t.Errorf("output not as expected.\nExpected %+v\n     Got %+v\n", val, gotVal)
		}
		t.Errorf("Key not found in output: %s\n", key)
	}

}

func TestJsonToBsonMapCast(t *testing.T) {

	inputData := map[string]string{
		"jstring":   "string",
		"jint":      "-1",
		"juint":     "2",
		"jobjectid": bson.NewObjectId().Hex(),
	}

	caster := NewMapCaster()
	caster.JsonInput()
	caster.BsonOutput()

	targetStruct := inputStruct{}
	outputMap := caster.Cast(inputData, &targetStruct)

	expectedOutput := map[string]interface{}{
		"bstring":   "string",
		"bint":      -1,
		"buint":     uint(2),
		"bobjectid": bson.ObjectIdHex(inputData["jobjectid"]),
	}

	for key, val := range expectedOutput {
		if gotVal, found := outputMap[key]; found == true {
			if gotVal == val {
				t.Log("Value matches:", key, val, gotVal)
				continue
			}
			t.Errorf("output not as expected.\nExpected %+v\n     Got %+v\n", val, gotVal)
		}
		t.Errorf("Key not found in output: %s\noutput:%+v\n", key, outputMap)
	}

}

func TestJsonToProtoMapCast(t *testing.T) {

	inputData := map[string]string{
		"jstring":   "string",
		"jint":      "-1",
		"juint":     "2",
		"jobjectid": bson.NewObjectId().Hex(),
	}

	caster := NewMapCaster()
	caster.JsonInput()
	caster.ProtoOutput()

	targetStruct := inputStruct{}
	outputMap := caster.Cast(inputData, &targetStruct)

	expectedOutput := map[string]interface{}{
		"pstring":   "string",
		"pint":      -1,
		"puint":     uint(2),
		"pobjectid": bson.ObjectIdHex(inputData["jobjectid"]),
	}

	for key, val := range expectedOutput {
		if gotVal, found := outputMap[key]; found == true {
			if gotVal == val {
				t.Log("Value matches:", key, val, gotVal)
				continue
			}
			t.Errorf("output not as expected.\nExpected %+v\n     Got %+v\n", val, gotVal)
		}
		t.Errorf("Key not found in output: %s\noutput:%+v\n", key, outputMap)
	}

}
