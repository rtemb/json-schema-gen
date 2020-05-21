package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/alecthomas/jsonschema"
	"github.com/rtemb/grpc-gateway-test/webhook_models/my_subscription"
)

var requestsForSchemas = []interface{}{
	my_subscription.Request_V1{Payload: my_subscription.Payload_V1{}},
	my_subscription.Request_V2{Payload: my_subscription.Payload_V2{}},
}

func main() {
	reflector := &jsonschema.Reflector{ExpandedStruct: true}
	for _, req := range requestsForSchemas {
		schema := reflector.Reflect(req)
		validationSchema, _ := json.MarshalIndent(schema, "", "  ")
		err := ioutil.WriteFile("docs/jsonschema/"+reflect.TypeOf(req).String()+".jsonschema", validationSchema, 777)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	}
}
