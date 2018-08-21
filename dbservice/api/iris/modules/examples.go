/*
@Time : 2018/8/21 11:19 
@Author : kenny zhu
@File : examples.go
@Software: GoLand
@Others:
*/
package modules

import (
	"github.com/kataras/iris"
	"github.com/kennyzhu/go-os/log"
	example "github.com/kennyzhu/go-os/dbservice/proto/example"
	"context"
	"encoding/json"
)

type examples struct{
	cl example.PreferencesService
}

// All are run in goroutine
func (s *examples) Preferences(ctx iris.Context) {
	action := ctx.Params().Get("action")

	switch action {
	case "GetPreference":
		s.getPrefernces( ctx )
	case "GetPreferencesList":
		s.getPreferencesList( ctx )
	default:
		ctx.JSON(iris.Map{
			"message": "Unknown action:" + action,
		})
	}
	log.Debug("Preferences done!")
}

func (s *examples) getPrefernces(ctx iris.Context) {
	log.Info("Received getPrefernces http request")

	user := ctx.URLParamInt32Default("user", 1)

	response, err := s.cl.GetPreference(context.TODO(), &example.PreferenceRequest{
		User: user,
	})

	if err != nil {
		ctx.StatusCode(500)
		ctx.JSON(iris.Map{
			"message": err,
		})
		return
	}

	ctx.StatusCode( int(response.ResultCode) )

	prefersJson,_ := json.Marshal( response.Prefer )
	b, _ := json.Marshal(map[string]string{
		"message": string(prefersJson[:]),
	})
	ctx.JSON(b)
	log.Info(prefersJson)
}

func (s *examples) getPreferencesList(ctx iris.Context) {
	log.Info("Received getPreferencesList http request")

	index := ctx.URLParamInt32Default("index", 1)
	limit := ctx.URLParamInt32Default("limit", 2)

	response, err := s.cl.GetPreferencesList(context.TODO(), &example.PreferencesListRequest{
		Index:  index ,
		Limit: limit,
	})
	if err != nil {
		ctx.StatusCode(500)
		ctx.JSON(iris.Map{
			"message": err,
		})
		return
	}

	ctx.StatusCode( int(response.ResultCode) )
	prefersJson,_ := json.Marshal( response.Prefers )
	b, _ := json.Marshal(map[string]string{
		"message": string(prefersJson[:]),
	})
	ctx.JSON(b)
	log.Info(prefersJson)
}
