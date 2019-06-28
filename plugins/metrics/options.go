/*
@Time : 2019/5/10 10:16 
@Author : kenny zhu
@File : options.go
@Software: GoLand
@Others:
*/
package metrics

type Options struct {
	Namespace     string
	Id            string
	Version    	  string
	MetaData      map[string]string
	Function      map[string]func() float64
}


func Id(n string) Option {
	return func(o *Options) {
		o.Id = n
	}
}

func Namespace(n string) Option {
	return func(o *Options) {
		o.Namespace = n
	}
}

func Version(n string) Option {
	return func(o *Options) {
		o.Version = n
	}
}

func MetaData(n map[string]string) Option {
	return func(o *Options) {
		o.MetaData = n
	}
}

func Function(n map[string]func() float64) Option {
	return func(o *Options) {
		o.Function = n
	}
}
