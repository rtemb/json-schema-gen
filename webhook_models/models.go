package webhook_models

type Envelope struct {
	Timestamp int    `json:"timestamp"`
	EventID   string `json:"event_id" jsonschema:"example=qweqwe-12314"`
}

type Payload interface{}

type Request struct {
	Envelope
	Payload Payload `json:"payload"`
}
