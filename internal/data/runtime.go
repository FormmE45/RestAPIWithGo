package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// implement a MarshalJSON method on the Runtime type so that it
// satisfies the json.Marshaler interface for custom usage of Json-encoded value for the type
func (runtime Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", runtime)
	//Need to be quoted for a valid Json string	because we want MarshalJSON to return a slice of byte of a string
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
