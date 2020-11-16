//Every handler runs in a goroutine.
package handler

import (
	"context"
	models2 "github.com/kennyzhu/go-os/src/dbservice/models"
	"github.com/kennyzhu/go-os/src/dbservice/proto/example"
	log2 "github.com/kennyzhu/go-os/src/log"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
// Preferences.GetPreference
func (e *Example) GetPreference(ctx context.Context, req *go_micro_srv_dbservice.PreferenceRequest, rsp *go_micro_srv_dbservice.PreferenceResponse) error {
	log2.Infof("Received Example.GetPreference request,id %v", req.User)

	//select db to request..
	//no sharding here..
	userId := req.User
	preferResult := &models2.Preferences{}
	err := models2.GetById(int64(userId), preferResult)
	if err != nil {
		rsp.ResultCode = 500
		log2.Error(err)
	} else {
		rsp.ResultCode = 200
	}

	// protobuf中所有对象都是指针....
	rsp.Prefer = new(go_micro_srv_dbservice.Preference)
	rsp.Prefer.User = int32(preferResult.User)
	rsp.Prefer.Name = preferResult.Name
	rsp.Prefer.Value = preferResult.Value
	return nil
}

// get all list.
func (e *Example) GetPreferencesList(ctx context.Context, req *go_micro_srv_dbservice.PreferencesListRequest, rsp *go_micro_srv_dbservice.PreferencesListResponse) error {
	log2.Info("Received Example.GetPreferencesList request with limits: ", req.Limit)

	pPreferList := make([]*models2.Preferences, 0)
	err := models2.Find(int(req.Limit), int(req.Index), &pPreferList)
	if err != nil {
		rsp.ResultCode = 500
		log2.Error(err)
	} else {
		rsp.ResultCode = 200
	}

	for _,prefer := range pPreferList {
		preference := &go_micro_srv_dbservice.Preference{}
		preference.User = int32(prefer.User)
		preference.Value = prefer.Value
		preference.Name = prefer.Name
		rsp.Prefers = append(rsp.Prefers, preference)
	}

	log2.Info("Example.GetPreferencesList result counts: %d", len(rsp.Prefers))
	return nil
}
