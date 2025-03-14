package memefish

import (
	"bytes"
	"fmt"

	"github.com/cloudspannerecosystem/memefish/token"
)

// MultiError is a list of errors occured on parsing.
//
// Note that Parse* methods returns this wrapped error even if the error is just one.
type MultiError []*Error

func (list MultiError) String() string {
	return list.Error()
}

// Error returns an error message.
//
// This message only shows the first error's message and other errors' messages are omitted.
// If you want to obtain all messages of errors at once, you can use FullError instead.
func (list MultiError) Error() string {
	switch len(list) {
	case 0:
		return "(0 errors)"
	case 1:
		return list[0].Error()
	case 2:
		return list[0].Error() + " (and 1 other error)"
	default:
		return fmt.Sprintf("%s (and %d other errors)", list[0].Error(), len(list)-1)
	}
}

// FullError returns a full error message.
func (list MultiError) FullError() string {
	var message bytes.Buffer
	for _, err := range list {
		fmt.Fprintln(&message, err.FullError())
	}
	return message.String()
}

// Error is an error occured on parsing.
type Error struct {
	Message  string
	Position *token.Position
}

func (e *Error) String() string {
	return e.Error()
}

// Error returns an error message.
func (e *Error) Error() string {
	return fmt.Sprintf("syntax error: %s: %s", e.Position, e.Message)
}

func (e *Error) FullError() string {
	var message bytes.Buffer
	fmt.Fprint(&message, e.Error())
	if e.Position.Source != "" {
		fmt.Fprintf(&message, "\n%s", e.Position.Source)
	}
	return message.String()
}
