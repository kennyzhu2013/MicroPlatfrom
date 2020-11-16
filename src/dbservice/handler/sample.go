/*
@Time : 2018/9/27 10:59 
@Author : kenny zhu
@File : sample.go
@Software: GoLand
@Others:
*/
package handler

import (
	models2 "github.com/kennyzhu/go-os/src/dbservice/models"
)

type Sample struct{}

// Call is a single request handler called via client.Call or the generated client code
// Preferences.GetPreference
func (e *Example) GetArticle(ctx context.Context, req *example.PreferenceRequest, rsp *example.PreferenceResponse) error {
	log.Infof("Received Example.GetPreference request,id %v", req.User)

	//select db to request..
	//no sharding here..
	userId := req.User
	preferResult := &models2.Preferences{}
	err := models2.GetById(int64(userId), preferResult)
	if err != nil {
		rsp.ResultCode = 500
		log.Error(err)
	} else {
		rsp.ResultCode = 200
	}

	// protobuf中所有对象都是指针....
	rsp.Prefer = new(example.Preference)
	rsp.Prefer.User = int32(preferResult.User)
	rsp.Prefer.Name = preferResult.Name
	rsp.Prefer.Value = preferResult.Value
	return nil
}
