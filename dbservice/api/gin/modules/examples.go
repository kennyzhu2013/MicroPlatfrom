/*
@Time : 2018/8/21 11:19 
@Author : kenny zhu
@File : examples.go
@Software: GoLand
@Others:
*/
package modules

import (
	"github.com/kennyzhu/go-os/log"
	example "github.com/kennyzhu/go-os/dbservice/proto/example"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

type examples struct{
	cl example.PreferencesService
}

// All are run in goroutine
func (s *examples) Preferences(ctx *gin.Context) {
	action := ctx.Param("action")

	switch action {
	case "/GetPreference":
		s.getPrefernces( ctx )
	case "/GetPreferencesList":
		s.getPreferencesList( ctx )
	default:
		ctx.JSON(404, map[string]string {
			"message": "Unknown action:" + action,
		})
	}
	log.Debug("Preferences done!")
}

func (s *examples) getPrefernces(ctx *gin.Context) {
	log.Info("Received getPrefernces http request")

	user,_ := strconv.Atoi ( ctx.DefaultQuery("user", "1") )

	response, err := s.cl.GetPreference(context.TODO(), &example.PreferenceRequest{
		User: int32( user ),
	})

	if err != nil {
		ctx.JSON(500, map[string]interface{}{
			"message": err,
		})
		return
	}

	prefersJson,_ := json.Marshal( response.Prefer )
	b, _ := json.Marshal(map[string]string{
		"message": string(prefersJson[:]),
	})
	ctx.JSON(int(response.ResultCode), b)
	log.Info(prefersJson)
}

func (s *examples) getPreferencesList(ctx *gin.Context) {
	log.Info("Received getPreferencesList http request")

	index,_ :=  strconv.Atoi ( ctx.Query("index") )
	limit,_ :=  strconv.Atoi ( ctx.Query("limit") )

	response, err := s.cl.GetPreferencesList(context.TODO(), &example.PreferencesListRequest{
		Index:  int32(index) ,
		Limit: int32(limit),
	})
	if err != nil {
		ctx.JSON(500, map[string]interface{}{
			"message": err,
		})
		return
	}

	prefersJson,_ := json.Marshal( response.Prefers )
	b, _ := json.Marshal(map[string]string{
		"message": string(prefersJson[:]),
	})
	ctx.JSON(int(response.ResultCode), b)
	log.Info(prefersJson)
}
