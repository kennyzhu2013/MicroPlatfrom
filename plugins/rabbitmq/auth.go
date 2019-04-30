/*
@Time : 2019/4/28 11:06 
@Author : kenny zhu
@File : auth.go
@Software: GoLand
@Others:
*/
package rabbitmq

// fot auth key.
type ExternalAuthentication struct {
}

func (auth *ExternalAuthentication) Mechanism() string {
	return "EXTERNAL"
}

func (auth *ExternalAuthentication) Response() string {
	return ""
}