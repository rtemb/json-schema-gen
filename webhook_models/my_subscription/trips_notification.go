package my_subscription

import "github.com/rtemb/grpc-gateway-test/webhook_models"

const Subscription = "MY_SUBSCRIPTION"

type Request_V1 struct {
	webhook_models.Envelope
	Payload Payload_V1 `json:"payload"`
}
type Payload_V1 struct {
	Name    string `json:"name"`
	EventID string `json:"event_id" jsonschema:"title=the name,description=The name of a friend,example=joe,example=lucy,default=alex"`
}

type Request_V2 struct {
	webhook_models.Envelope
	Payload Payload_V2 `json:"payload"`
}
type Payload_V2 struct {
	Name     string `json:"name"`
	Source   string `json:"source"`
	DateTime string `json:"date_time"`
}
