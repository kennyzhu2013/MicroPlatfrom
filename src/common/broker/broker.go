/*
@Time : 2019/6/5 16:28
@Author : kenny zhu
@File : broker
@Software: GoLand
@Others:
*/
package broker


// Broker is an interface used for asynchronous messaging.
type Broker interface {
	Options() Options
	Address() string
	Connect() error
	Disconnect() error
	Init(...Option) error
	Publish(string, *Message, ...PublishOption) error
	Subscribe(string, Handler, ...SubscribeOption) (Subscriber, error)
	String() string
}

// Handler is used to process messages via a subscription of a topic.
type Handler func(Publication) error

type Message struct {
	Header map[string]string
	Body   []byte
}

// Publication is given to a subscription handler for processing
type Publication interface {
	Topic() string
	Message() *Message
	Ack() error
}

// Subscriber is a convenience return type for the Subscribe method
type Subscriber interface {
	Options() SubscribeOptions
	Topic() string
	Unsubscribe() error
}