# mapcast
Convert `map[string]string` into `map[string]interface` using a reference struct to sniff types. Optionally can expect the struct's json tag names. Optionally can returned data to `bson` tag name. 

# WHY?

In order to use query string values as values in mongo db calls.
That's my use case.

## Cast
Casts `map[string]string` referencing any `struct` to `map[string]interface`

## CastViaJson
Casts `map[string]string` referencing any `struct` to `map[string]interface` expecting json name in input

Respects the "-" flag.

## CastViaJsonToBson
Casts `map[string]string` referencing any `struct` to `map[string]interface` expecting json name in input and returning bson field names in map

## CastViaProtobufToBson
Casts `map[string]string` referencing any `struct` to `map[string]interface` expecting protobuf name in input and returning bson field names in map

## Examples

    type myStruct struct {
        Field int `json:"input_name" bson:"output_name"`
        Hidden float32 `json:"-" bson:"hidden_field"`
    }
    
    myInput := map[string]string{"input_name": "345"}
    
    Cast(myInput, myStruct{}) 
    // returns map["Field" : 345]
    
    CastViaJson(myInput, myStruct{}) 
    // returns map["input_name" : 345]
    
    CastViaJsonToBson(myInput, myStruct{}) 
    // returns map["output_name" : 345]

### Supported Types

It can convert to these types so far (againg ObjectId is here since my use case is mongo related)

- bool
- string
- int
- int8
- int16
- int32
- uint
- uint8
- uint16
- uint32
- float64
- float128
- bson.ObjectId
