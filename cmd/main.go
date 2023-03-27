package main

import (
	"fmt"
	googleimage "github.com/fsuke1985/jacreen/core/types"
	"encoding/json"
	io "io/ioutil"
	"os"
	"net/http"
)

func main() {
	var x *googleimage.GoogleImage

	res, err := http.Get("http://apiserver/api/endpoint/data.json")
	if err != nil {
		os.Exit(1)
	}
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		os.Exit(1)
	}

	json.Unmarshal(raw, &x)

	for _, i := range x.Error.Errors {
		fmt.Println(x.Error.Code)
		fmt.Println(i.Message)
	}
}