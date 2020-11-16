package main

import (
	"github.com/kennyzhu/go-os/src/dbservice/api/conf"
	modules2 "github.com/kennyzhu/go-os/src/dbservice/api/modules"
	log2 "github.com/kennyzhu/go-os/src/log"
	"github.com/micro/go-micro"
)

/*
type Preferences struct {
	Client example.PreferencesService
}

//process all message: Say/PreferenceList?....
func (s *Preferences) PreferenceList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Preferences.PreferenceList API request")

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
	response, err := s.Client.GetPreferencesList(ctx, &example.PreferencesListRequest{
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
}*/

// usage: curl "http://localhost:8002/dbservice/Preferences/PreferenceList?limit=2&index=1".
//       or http.Post("http://localhost:8002/dbservice/Preferences/PreferenceList", "application/protobuf", bytes.NewReader(req))
//       or json:
func main() {
	conf.Init("./conf/api.json")
	service := micro.NewService(
		micro.Name(conf.AppConf.ApiName), // eg: "go.micro.api.dbservice".
	)

	// parse command line flags
	service.Init()
	modules2.Init( service.Server(), service.Client() )
	/*
	service.Server().Handle(
		service.Server().NewHandler(
			&Preferences{Client: example.NewPreferencesService(AppConf.SrvName, service.Client())},
		),
	)*/

	if err := service.Run(); err != nil {
		log2.Fatal(err)
	}
}
