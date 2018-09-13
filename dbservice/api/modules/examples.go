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
	api "github.com/kennyzhu/go-os/dbservice/proto/api"
	"github.com/micro/go-micro/errors"
	"strconv"
	. "github.com/kennyzhu/go-os/dbservice/api/conf"
)

// examples is service name
type examples struct{
	cl example.PreferencesService
}

func (s *examples) Preference(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Preference http request")

	user, ok := req.Get["index"]
	if !ok || len(user.Values) == 0 {
		return errors.BadRequest(AppConf.SrvName, "Index cannot be blank")
	}
	userInt,_ := strconv.Atoi( user.Values[0] )

	response, err := s.cl.GetPreference(context.TODO(), &example.PreferenceRequest{
		User: int32( userInt ),
	})

	if err != nil {
		log.Error(err)
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
func (s *examples) PreferenceList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received PreferenceList API request")

	index, ok := req.Get["index"]
	if !ok || len(index.Values) == 0 {
		return errors.BadRequest(AppConf.SrvName, "Index cannot be blank")
	}

	limit, ok := req.Get["limit"]
	if !ok || len(limit.Values) == 0 {
		return errors.BadRequest(AppConf.SrvName, "Index cannot be blank")
	}

	indexInt,_ := strconv.Atoi( index.Values[0] )
	limitInt,_ := strconv.Atoi( limit.Values[0] )
	response, err := s.cl.GetPreferencesList(ctx, &example.PreferencesListRequest{
		Index: int32( indexInt ),
		Limit: int32( limitInt ),
	})
	if err != nil {
		log.Error(err)
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
