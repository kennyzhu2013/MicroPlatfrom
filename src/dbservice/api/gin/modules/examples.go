/*
@Time : 2018/8/21 11:19 
@Author : kenny zhu
@File : examples.go
@Software: GoLand
@Others:
*/
package modules

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kennyzhu/go-os/src/dbservice/proto/example"
	log2 "github.com/kennyzhu/go-os/src/log"
	"strconv"
)

type examples struct{
	cl go_micro_srv_dbservice.PreferencesService
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
	log2.Debug("Preferences done!")
}

func (s *examples) getPrefernces(ctx *gin.Context) {
	log2.Info("Received getPrefernces http request")

	user,_ := strconv.Atoi ( ctx.DefaultQuery("user", "1") )

	response, err := s.cl.GetPreference(context.TODO(), &go_micro_srv_dbservice.PreferenceRequest{
		User: int32( user ),
	})

	if err != nil {
		ctx.JSON(500, map[string]string{
			"message": err.Error(),
		})
		log2.Error(err)
		return
	}

	/*
	prefersJson,_ := json.Marshal( response.Prefer )
	ctx.JSON(int(response.ResultCode), map[string]string{
		"message": string(prefersJson[:]),
	})*/
	ctx.JSON(int(response.ResultCode), gin.H{
		"message": response.Prefer,
	})

	log2.Info("getPrefernces End:")
	// log.Info(b)
}

func (s *examples) getPreferencesList(ctx *gin.Context) {
	log2.Info("Received getPreferencesList http request")

	index,_ :=  strconv.Atoi ( ctx.Query("index") )
	limit,_ :=  strconv.Atoi ( ctx.Query("limit") )

	response, err := s.cl.GetPreferencesList(context.TODO(), &go_micro_srv_dbservice.PreferencesListRequest{
		Index:  int32(index) ,
		Limit: int32(limit),
	})
	if err != nil {
		ctx.JSON(500, map[string]string{
			"message": err.Error(),
		})
		log2.Error(err)
		return
	}

	prefersJson,_ := json.Marshal( response.Prefers )
	/*
	b, _ := json.Marshal(map[string]string{
		"message": string(prefersJson[:]),
	})*/
	ctx.JSON(int(response.ResultCode), map[string]string{
		"message": string(prefersJson[:]),
	})
	log2.Info("getPreferencesList End:")
	// log.Info(b)
}
