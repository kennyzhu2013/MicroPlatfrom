/*
@Time : 2019/4/28 10:44 
@Author : kenny zhu
@File : context.go
@Software: GoLand
@Others:
*/
package rabbitmq

import "github.com/micro/go-micro/broker"
import "context"

// setSubscribeOption returns a function to setup a context with given value
func setSubscribeOption(k, v interface{}) broker.SubscribeOption {
	return func(o *broker.SubscribeOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

// setBrokerOption returns a function to setup a context with given value
func setBrokerOption(k, v interface{}) broker.Option {
	return func(o *broker.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

// setPublishOption returns a function to setup a context with given value
func setPublishOption(k, v interface{}) broker.PublishOption {
	return func(o *broker.PublishOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}
