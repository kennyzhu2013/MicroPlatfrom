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
	"github.com/kennyzhu/go-os/src/dbservice/api/conf"
	"github.com/kennyzhu/go-os/src/dbservice/proto/api"
	"github.com/kennyzhu/go-os/src/dbservice/proto/example"
	log2 "github.com/kennyzhu/go-os/src/log"
	"github.com/micro/go-micro/errors"
	"strconv"
)

// examples is service name
type examples struct{
	cl go_micro_srv_dbservice.PreferencesService
}

func (s *examples) Preference(ctx context.Context, req *go_micro_api.Request, rsp *go_micro_api.Response) error {
	log2.Info("Received Preference http request")

	user, ok := req.Get["index"]
	if !ok || len(user.Values) == 0 {
		return errors.BadRequest(conf.AppConf.SrvName, "Index cannot be blank")
	}
	userInt,_ := strconv.Atoi( user.Values[0] )

	response, err := s.cl.GetPreference(context.TODO(), &go_micro_srv_dbservice.PreferenceRequest{
		User: int32( userInt ),
	})

	if err != nil {
		log2.Error(err)
		rsp.StatusCode = 500
		return err
	}

	rsp.StatusCode = response.ResultCode
	prefersJson,_ := json.Marshal( response.Prefer )
	b, _ := json.Marshal(map[string]string{
		"message": string(prefersJson[:]),
	})
	rsp.Body = string(b)
	return nil
}

// process all message: Say/PreferenceList?....
func (s *examples) PreferenceList(ctx context.Context, req *go_micro_api.Request, rsp *go_micro_api.Response) error {
	log2.Info("Received PreferenceList API request")

	index, ok := req.Get["index"]
	if !ok || len(index.Values) == 0 {
		return errors.BadRequest(conf.AppConf.SrvName, "Index cannot be blank")
	}

	limit, ok := req.Get["limit"]
	if !ok || len(limit.Values) == 0 {
		return errors.BadRequest(conf.AppConf.SrvName, "Index cannot be blank")
	}

	indexInt,_ := strconv.Atoi( index.Values[0] )
	limitInt,_ := strconv.Atoi( limit.Values[0] )
	response, err := s.cl.GetPreferencesList(ctx, &go_micro_srv_dbservice.PreferencesListRequest{
		Index: int32( indexInt ),
		Limit: int32( limitInt ),
	})
	if err != nil {
		log2.Error(err)
		rsp.StatusCode = 500
		return err
	}

	rsp.StatusCode = response.ResultCode
	prefersJson,_ := json.Marshal( response.Prefers )
	b, _ := json.Marshal(map[string]string{
		"message": string(prefersJson[:]),
	})

	rsp.Body = string(b)
	return nil
}
