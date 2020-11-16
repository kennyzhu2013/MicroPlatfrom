/*
@Time : 2018/8/21 11:05 
@Author : kenny zhu
@File : modules.go
@Software: GoLand
@Others:
*/
package modules

import (
	"github.com/kataras/iris"
	log2 "github.com/kennyzhu/go-os/src/log"

	"encoding/json"
	"io/ioutil"

	// "github.com/micro/go-micro/codec/protorpc"
	"fmt"
	proto "github.com/micro/go-micro/server/debug/proto"
)

var Modules struct{
	App *iris.Application

}

// self init...
func init() {
	Modules.App = iris.Default()
}

// {"method":"Debug.Health","params":[null],"id":85}..
type serverRequest struct {
	Method string           `json:"method"`
	Params *json.RawMessage `json:"params"`
	ID     *json.RawMessage `json:"id"`
}

type serverResponse struct {
	ID     *json.RawMessage `json:"id"`
	Result interface{}      `json:"result"`
	Error  interface{}      `json:"error"`
}

// json内容返回...
// curl -d 'service=go.micro.srv.greeter' \
//     -d 'method=Debug.Health' \
//     -d 'request={"name": "John"}' \
//     http://localhost:8080/rpc
//    只支持export MICRO_PROXY_ADDRESS=0.0.0.0:8002的jsonrpc形势下的health检查...
func NoModules(ctx iris.Context) {
	log2.Info("Received NoModules API request")
	b, _ := ioutil.ReadAll(ctx.Request().Body)
	// jrpc.NewCodec()
	var body serverRequest
	if err := json.Unmarshal(b, &body); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println( string(b) )
	fmt.Println( body )

	rsp := &serverResponse{}
	rsp.ID = body.ID

	result := &proto.HealthResponse{}
	result.Status = "ok"
	rsp.Result = result

	prefersJson,_ := json.Marshal( rsp )
	ctx.Write(prefersJson)

	// ctx.WriteString("{Status:\"ok\"}")
	/*
	ctx.JSON(map[string]string{
		"Status": "ok",
	})*/
}

//anything to do..
//run here not go routine...