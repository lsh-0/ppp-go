package utils

// something of a dumping ground for 'convenience' functions.

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func Pprint(val interface{}) {

	switch valtype := val.(type) {
	// list of somethings
	case []interface{}:
		Pprint("list of somethings")
		for i, v := range valtype {
			fmt.Print("index:", i, ": ")
			Pprint(v)
		}

	// map of somethings
	case map[string]interface{}:
		Pprint("maps of somethings")
		for k, v := range valtype {
			fmt.Print("key:", k, ": ")
			Pprint(v)
		}

	case string:
		fmt.Println(val)

	default:
		fmt.Println("unknown:", val)
	}

}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// coerce given `data` in to a JSON string
func ToJSON(data interface{}) string {
	b, err := json.Marshal(data)
	Check(err)
	return string(b[:])
}

// coerce given `bytes` into an instance of `T`
func FromJSON[T any](bytes []byte) T {
	ti := new(T)
	json.Unmarshal(bytes, &ti)
	return *ti
}

// read the JSON in `filename` into an instance of `T`
func ReadJSON[T any](filename string) T {
	bytes, err := os.ReadFile(filename)
	Check(err)
	return FromJSON[T](bytes)
}

// returns `true` if given `filename` exists.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}
