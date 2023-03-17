package errors

import "errors"

var (
	QueueEmptyError error
)

func init() {
	QueueEmptyError = errors.New("queue is empty")
}
