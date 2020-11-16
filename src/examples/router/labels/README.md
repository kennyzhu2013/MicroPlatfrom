# Label Based Filtered

This is an example of how to filter on labels. In the example we're injecting a routing label 
but in a real world environment the label may be set in a header from a client app/mobile app.

## Running Example

1. Start the routing-srv in github.com/micro/router-srv
2. Start the greeter-srv in github.com/micro/examples/greeter/srv
3. Run the query 
	micro query go.micro.srv.router Label.Update '{"label": {"id": "1", "service": "go.micro.srv.greeter", "version": "1.0.0", "key": "greeter", "value": "one", "weight": 50} }'
4. Run the example

You'll see that 50% of requests are matched to the service as we've set this label to be applied half of the time.


