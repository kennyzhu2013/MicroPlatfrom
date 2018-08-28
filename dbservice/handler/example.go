//Every handler runs in a goroutine.
package handler

import (
	"context"

	"github.com/kennyzhu/go-os/log"
	"github.com/kennyzhu/go-os/dbservice/models"
	example "github.com/kennyzhu/go-os/dbservice/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
// Preferences.GetPreference
func (e *Example) GetPreference(ctx context.Context, req *example.PreferenceRequest, rsp *example.PreferenceResponse) error {
	log.Infof("Received Example.GetPreference request,id %v", req.User)

	//select db to request..
	//no sharding here..
	userId := req.User
	preferResult := &models.Preferences{}
	err := models.GetById(int64(userId), preferResult)
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

// get all list.
func (e *Example) GetPreferencesList(ctx context.Context, req *example.PreferencesListRequest, rsp * example.PreferencesListResponse) error {
	log.Info("Received Example.GetPreferencesList request with limits: ", req.Limit)

	pPreferList := make([]*models.Preferences, 0)
	err := models.Find(int(req.Limit), int(req.Index), &pPreferList)
	if err != nil {
		rsp.ResultCode = 500
		log.Error(err)
	} else {
		rsp.ResultCode = 200
	}

	for _,prefer := range pPreferList {
		preference := &example.Preference{}
		preference.User = int32(prefer.User)
		preference.Value = prefer.Value
		preference.Name = prefer.Name
		rsp.Prefers = append(rsp.Prefers, preference)
	}

	log.Info("Example.GetPreferencesList result counts: %d", len(rsp.Prefers))
	return nil
}
