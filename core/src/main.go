package main

import (
	"fmt"
	googleimage "github.com/fsuke1985/jacreen/core/types"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	var x *googleimage.GoogleImage
	var j googleimage.GoogleImage

	raw, err := ioutil.ReadFile("../data.json")
	if err != nil {
		os.Exit(1)
	}

	json.Unmarshal(raw, &x)

	kJson := reflect.TypeOf(j)
	// vJson := reflect.ValueOf(x)

	for i := 0; i < kJson.NumField(); i++ {
		// f := kJson.Field(i)
		v, b := kJson.FieldByName("Code")
		
		fmt.Println(v.Code)
	}
}