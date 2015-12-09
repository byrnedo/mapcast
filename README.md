# mapcast
Convert map[string]string into map[string]interface using a reference struct. Optionally expecting json field name. Optionally map to bson field name. 

## Cast
Casts `map[string]string` to `map[string]interface`

## CastViaJson
Casts `map[string]string` to `map[string]interface` expecting json name in input

## CastViaJsonToBson
Casts `map[string]string` to `map[string]interface` expecting json name in input and returning bson field names in map
