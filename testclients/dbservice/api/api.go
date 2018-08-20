package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"
	example "github.com/kennyzhu/go-os/dbservice/proto/example"
)

//default api, refer examples/api/default/default.go
func main() {

	//getParams := make(map[string]*api.Pair)
	//indexValues := []string{"1", }
	//valuePair := api.Pair{Key: "", Values:indexValues, }
	//getParams["index"] = &api.Pair{ Key: "", Values:[]string{"1", },}
	//protobuf must define itself....
	/*
	req, err := proto.Marshal(&example.PreferencesListRequest{ Index:1, Limit:2, })
	if err != nil {
		fmt.Println(err)
		return
	}*/
	//for http...
	//"application/json"
	//r, err := http.Get("http://localhost:8002/dbservice/Preferences/PreferenceList", "application/protobuf", bytes.NewReader(req))
	r, err := http.Get("http://localhost:8002/dbservice/Preferences/PreferenceList?limit=2&index=1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(b)
	//return json...
	fmt.Println(r.StatusCode)
	var body map[string]interface{}
	if err := json.Unmarshal(b, &body); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(body["message"])
	rsp := make([]example.Preference, 0)
	//var rsp example.Preference[]
	if err := json.Unmarshal( []byte( body["message"].(string) ) , &rsp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}
