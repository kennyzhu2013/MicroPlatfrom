package web

import "io/ioutil"

var (
	layoutTemplate = ``
	indexTemplate = ``
	callTemplate  = ``
	registryTemplate = ``

	serviceTemplate = ``
	cliTemplate     = ``
)

// init template here....
func InitTemplateHandler(){
	bBuff,_ := ioutil.ReadFile("layout.tmpl")
	layoutTemplate = string(bBuff)

	bBuff,_ = ioutil.ReadFile("index.tmpl")
	indexTemplate = string(bBuff)

	bBuff,_ = ioutil.ReadFile("call.tmpl")
	callTemplate = string(bBuff)

	bBuff,_ = ioutil.ReadFile("registry.tmpl")
	registryTemplate = string(bBuff)

	bBuff,_ = ioutil.ReadFile("service.tmpl")
	serviceTemplate = string(bBuff)

	bBuff,_ = ioutil.ReadFile("cli.tmpl")
	cliTemplate = string(bBuff)
}
