// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/trace/v2/trace_service.proto

package envoy_service_trace_v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on StreamTracesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamTracesResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StreamTracesResponseValidationError is the validation error returned by
// StreamTracesResponse.Validate if the designated constraints aren't met.
type StreamTracesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamTracesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamTracesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamTracesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamTracesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamTracesResponseValidationError) ErrorName() string {
	return "StreamTracesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StreamTracesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamTracesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamTracesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamTracesResponseValidationError{}

// Validate checks the field values on StreamTracesMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamTracesMessage) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIdentifier()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamTracesMessageValidationError{
				field:  "Identifier",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSpans() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamTracesMessageValidationError{
					field:  fmt.Sprintf("Spans[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// StreamTracesMessageValidationError is the validation error returned by
// StreamTracesMessage.Validate if the designated constraints aren't met.
type StreamTracesMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamTracesMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamTracesMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamTracesMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamTracesMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamTracesMessageValidationError) ErrorName() string {
	return "StreamTracesMessageValidationError"
}

// Error satisfies the builtin error interface
func (e StreamTracesMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamTracesMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamTracesMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamTracesMessageValidationError{}

// Validate checks the field values on StreamTracesMessage_Identifier with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamTracesMessage_Identifier) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetNode() == nil {
		return StreamTracesMessage_IdentifierValidationError{
			field:  "Node",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamTracesMessage_IdentifierValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// StreamTracesMessage_IdentifierValidationError is the validation error
// returned by StreamTracesMessage_Identifier.Validate if the designated
// constraints aren't met.
type StreamTracesMessage_IdentifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamTracesMessage_IdentifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamTracesMessage_IdentifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamTracesMessage_IdentifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamTracesMessage_IdentifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamTracesMessage_IdentifierValidationError) ErrorName() string {
	return "StreamTracesMessage_IdentifierValidationError"
}

// Error satisfies the builtin error interface
func (e StreamTracesMessage_IdentifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamTracesMessage_Identifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamTracesMessage_IdentifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamTracesMessage_IdentifierValidationError{}
