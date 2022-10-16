package context

import (
	"github.com/loopholelabs/scale-go/runtime/generated"
)

// Request is the HTTP Request object
type Request struct {
	value *generated.Request
}

// Request returns the Request object for the Context
func (ctx *Context) Request() *Request {
	return &Request{value: ctx.generated.Request}
}

// Method returns the method of the request
func (req *Request) Method() string {
	return req.value.Method
}

// SetMethod sets the method of the request
func (req *Request) SetMethod(method string) string {
	req.value.Method = method
	return req.value.Method
}

// RemoteIP returns the remote IP of the request
func (req *Request) RemoteIP() string {
	return req.value.Ip
}

func (req *Request) Body() []byte {
	return req.value.Body
}

func (req *Request) SetBody(body string) {
	req.value.Body = []byte(body)
}

func (req *Request) SetBodyBytes(body []byte) {
	req.value.Body = body
}

// RequestHeaders is are the headers in the request
type RequestHeaders struct {
	value generated.RequestHeadersMap
}

// Headers returns the headers of the request
func (req *Request) Headers() *RequestHeaders {
	return &RequestHeaders{
		value: req.value.Headers,
	}
}

// Get returns the value of the header
func (h *RequestHeaders) Get(k string) []string {
	v := h.value[k]
	if v == nil {
		return nil
	}
	return v.Value
}

// Set sets the value of the header
func (h *RequestHeaders) Set(k string, v []string) {
	h.value[k] = &generated.StringList{
		Value: v,
	}
}
