package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

var logc logger.Logger

func init() {
	logc = logger.NewDefaultLogger()
}

type HttpMethod string

// http methods
const (
	Get HttpMethod = http.MethodGet
)

type HttpV1 struct {
	Method HttpMethod
}

func (h *HttpV1) GetMethod() string {
	return string(h.Method)
}

func (h *HttpV1) SetMethod(method string) error {
	method1 := strings.ToLower(method)

	switch {
	case method1 == "get":
		h.Method = Get
	default:
		return fmt.Errorf("invalid http method: %v, please select a valid http method", method)
	}

	return nil
}

func (*HttpV1) GetMethods() []HttpMethod {
	return []HttpMethod{Get}
}

type Request struct {
	Method HttpMethod
	URL    string
}

type Response struct {
	ErrorMsg string
	Body     []byte
	Headers  map[string][]string
}

func (h *HttpV1) SendRequest(req Request) Response {
	fmt.Println("req", req)

	var ir io.Reader

	if req.Method == Get {
		ir = http.NoBody
	}

	httpReq, err := http.NewRequest(string(req.Method), req.URL, ir)
	if err != nil {
		logc.Error("failed to create request reason: " + err.Error())
		return Response{
			ErrorMsg: err.Error(),
		}
	}

	client := &http.Client{}

	rawResp, err := client.Do(httpReq)
	if err != nil {
		logc.Error("failed to make the request reason: " + err.Error())

		return Response{
			ErrorMsg: err.Error(),
		}
	}

	logc.Info("received response")

	buf := bytes.Buffer{}

	_, err = buf.ReadFrom(rawResp.Body)
	if err != nil {
		logc.Error("failed to read from response reason: " + err.Error())

		return Response{
			ErrorMsg: err.Error(),
		}
	}

	logc.Info("response data: " + string(buf.Bytes()))

	return Response{
		Body:    buf.Bytes(),
		Headers: rawResp.Header.Clone(),
	}
}
