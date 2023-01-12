package errors

import "fmt"

// AddCode adds a error with a code to parent error.
func AddCode(parent error, code string) error {
	return &CodeErr{Parent: parent, Code: code}
}

// AddCode adds a error with a code and message to parent error.
func AddCodeWithMessage(parent error, code, message string) error {
	return &CodeErr{Parent: parent, Code: code, Msg: message}
}

// AddCodeWithrags adds a error with a code, message and arguments to parent error.
func AddCodeWithMessageAndArgs(parent error, code, message string, args map[string]interface{}) error {
	return &CodeErr{Parent: parent, Code: code, Msg: message, Args: &args}
}

// CodeErr is a error that includes a Code.
type CodeErr struct {
	Parent error
	Msg    string
	Code   string
	Args   *map[string]interface{}
}

func (c *CodeErr) Error() string {
	msg := fmt.Sprintf("code %s", c.Code)
	if c.Msg != "" {
		msg = c.Msg
	}

	if c.Parent != nil {
		return fmt.Sprintf("%s; %s", msg, c.Parent)
	}

	return msg
}

func (c *CodeErr) Unwrap() error {
	return c.Parent
}
