package errors

import "errors"

var (
	QueueEmptyError     error
	NoLegalWSConnection error
	ClientParseError    error
)

func init() {
	QueueEmptyError = errors.New("queue is empty")
	NoLegalWSConnection = errors.New("no legal websocket connection")
	ClientParseError = errors.New("value of sync.Map parse struct Client error")
}
