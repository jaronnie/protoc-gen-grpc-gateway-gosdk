// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Request allows for building up a request to a server in a chained fashion.
// Any errors are stored until the end of your call, so you only have to
// check once.
type Request struct {
	c *RESTClient

	verb string

	subPath string

	params string

	// output
	err  error
	body io.Reader
}

func NewRequest(c *RESTClient) *Request {

	r := &Request{
		c: c,
	}
	return r
}

func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

type PathParam struct {
	Name  string
	Value interface{}
}

// SubPath set subPath
// e.g. /api/v1/credential/init
func (r *Request) SubPath(subPath string, args ...PathParam) *Request {
	for _, v := range args {
		subPath = strings.ReplaceAll(subPath, "{"+v.Name+"}", cast.ToString(v.Value))
	}
	r.subPath = r.c.gatewayPrefix + subPath

	if r.c.disableGateway {
		i := strings.Index(r.subPath, "/api")
		r.subPath = r.subPath[i:]
	}

	return r
}

type QueryParam struct {
	Name  string
	Value interface{}
}

func (r *Request) Params(args ...QueryParam) *Request {
	if len(args) == 0 {
		return r
	}
	queryParams := "?"
	for i, v := range args {
		if cast.ToString(v.Value) == "" {
			continue
		}
		if i == len(args)-1 {
			queryParams += fmt.Sprintf("%s=%s", v.Name, cast.ToString(v.Value))
		} else {
			queryParams += fmt.Sprintf("%s=%s&", v.Name, cast.ToString(v.Value))
		}
	}
	r.params = queryParams
	return r
}

// defaultUrl get default url for common request
func (r *Request) defaultUrl() (string, error) {
	if r.c.protocol == "" || r.c.addr == "" || r.c.port == "" {
		return "", errors.New("invalid url, you may not login")
	}

	return fmt.Sprintf("%s://%s:%s", r.c.protocol, r.c.addr, r.c.port+r.subPath+r.params), nil
}

// WSUrl get WS url for request
func (r *Request) wsUrl() (string, error) {
	if r.c.protocol == "" || r.c.addr == "" || r.c.port == "" {
		return "", errors.New("invalid url, you may not login")
	}

	// upgrade http to websocket proto
	if r.c.protocol == "https" {
		r.c.protocol = "wss"
	} else {
		r.c.protocol = "ws"
	}

	return fmt.Sprintf("%s://%s:%s", r.c.protocol, r.c.addr, r.c.port+r.subPath+r.params), nil
}

// Body makes the request use obj as the body. Optional.
// If obj is a string, try to read a file of that name.
// If obj is a []byte, send it directly.
// default marshal it
func (r *Request) Body(obj interface{}) *Request {
	if r.err != nil {
		return r
	}

	switch t := obj.(type) {
	case string:
		r.body = bytes.NewReader([]byte(t))
	case []byte:
		r.body = bytes.NewReader(t)
	default:
		data, err := json.Marshal(obj)
		if err != nil {
			r.err = err
			return r
		}
		r.body = bytes.NewReader(data)
	}
	return r
}

// Result contains the result of calling Request.Do().
type Result struct {
	body       []byte
	err        error
	statusCode int
	status     string
}

// Do format and executes the request. Returns a Result object for easy response
// processing.
//
// Error type:
// http.Client.Do errors are returned directly.
func (r *Request) Do(ctx context.Context) Result {
	url, err := r.defaultUrl()
	if err != nil {
		return Result{err: err}
	}

	request, err := http.NewRequestWithContext(ctx, r.verb, url, r.body)
	if err != nil {
		return Result{err: err}
	}

	request.Header = r.c.headers

	if r.c.client == nil {
		r.c.client = http.DefaultClient
	}

	if r.c.retryTimes == 0 {
		r.c.retryTimes = 1
	}

	var rawResp *http.Response
	// if meet error, retry times that you set
	for k := 0; k < r.c.retryTimes; k++ {
		rawResp, err = doRequest(r.c.client, request)
		if err != nil {
			// sleep retry delay
			time.Sleep(r.c.retryDelay)
			continue
		}
		break
	}

	if err != nil {
		return Result{err: err}
	}

	if rawResp == nil {
		return Result{err: errors.New("http response is nil")}
	}

	defer rawResp.Body.Close()

	if rawResp.StatusCode != 200 {
		return Result{err: errors.Errorf("unhealthy status code: [%d], status message: [%s]", rawResp.StatusCode, rawResp.Status)}
	}

	data, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return Result{err: err, statusCode: rawResp.StatusCode, status: rawResp.Status}
	}
	return Result{
		body:       data,
		err:        nil,
		statusCode: rawResp.StatusCode,
		status:     rawResp.Status,
	}
}

// DoUpload format and executes the upload request. Returns a Result object for easy response
// processing.
//
// Error type:
// http.Client.Do errors are returned directly.
func (r *Request) DoUpload(ctx context.Context, fieldName string, filename string, filedata []byte) Result {
	url, err := r.defaultUrl()
	if err != nil {
		return Result{err: err}
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	part, err := writer.CreateFormFile(fieldName, filename)
	if err != nil {
		return Result{err: err}
	}
	_, err = io.Copy(part, bytes.NewReader(filedata))
	if err != nil {
		return Result{err: err}
	}
	err = writer.Close()
	if err != nil {
		return Result{err: err}
	}
	request, err := http.NewRequestWithContext(ctx, r.verb, url, payload)
	if err != nil {
		return Result{err: err}
	}

	headers := r.c.headers
	headers.Set("Content-Type", writer.FormDataContentType())
	request.Header = headers

	if r.c.client == nil {
		r.c.client = http.DefaultClient
	}

	if r.c.retryTimes == 0 {
		r.c.retryTimes = 1
	}

	var rawResp *http.Response
	// if meet error, retry times that you set
	for k := 0; k < r.c.retryTimes; k++ {
		rawResp, err = doRequest(r.c.client, request)
		if err != nil {
			// sleep retry delay
			time.Sleep(r.c.retryDelay)
			continue
		}
		break
	}

	if err != nil {
		return Result{err: err}
	}

	if rawResp == nil {
		return Result{err: errors.New("http response is nil")}
	}

	defer rawResp.Body.Close()

	data, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return Result{err: err, statusCode: rawResp.StatusCode, status: rawResp.Status}
	}

	if rawResp.StatusCode != 200 {
		return Result{err: errors.Errorf("unhealthy status code: [%d], status message: [%s]", rawResp.StatusCode, rawResp.Status), body: data}
	}

	return Result{
		body:       data,
		err:        nil,
		statusCode: rawResp.StatusCode,
		status:     rawResp.Status,
	}
}

func (r *Request) WsConn(ctx context.Context) (*websocket.Conn, *http.Response, error) {
	url, err := r.wsUrl()
	if err != nil {
		return nil, nil, err
	}
	return websocket.DefaultDialer.DialContext(ctx, url, r.c.headers)
}

func doRequest(client *http.Client, request *http.Request) (*http.Response, error) {
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("response is nil")
	}
	return res, nil
}

// Into stores the result into obj, if possible. If obj is nil it is ignored.
func (r Result) Into(obj interface{}, isWarpHttpResponse bool) error {
	if r.err != nil {
		return r.err
	}
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("object is not a ptr")
	}

	j, err := simplejson.NewJson(r.body)
	if err != nil {
		return err
	}

	// parse response data
	// code message data
	var marshalJSON []byte
	if isWarpHttpResponse {
		code, err := j.Get("code").Int()
		if err != nil {
			return err
		}
		if code != http.StatusOK {
			message, _ := j.Get("message").String()
			return fmt.Errorf(message)
		}
		data := j.Get("data")
		data.Del("@type") // 适配 grpc 存在的 @type 字段
		marshalJSON, err = data.MarshalJSON()
		if err != nil {
			return err
		}
	} else {
		marshalJSON, err = j.MarshalJSON()
		if err != nil {
			return err
		}
	}

	switch obj.(type) {
	case proto.Message:
		err = protojson.Unmarshal([]byte(marshalJSON), obj.(proto.Message))
	default:
		err = json.Unmarshal(marshalJSON, &obj)
	}

	if err != nil {
		return err
	}

	return nil
}

// StatusCode returns the HTTP status code of the request. (Only valid if no
// error was returned.)
func (r Result) StatusCode() int {
	return r.statusCode
}

// Stream proto Stream way return io.ReadCloser
func (r *Request) Stream(ctx context.Context) (io.ReadCloser, error) {
	url, err := r.defaultUrl()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, r.verb, url, r.body)
	if err != nil {
		return nil, err
	}

	request.Header = r.c.headers

	if r.c.client == nil {
		r.c.client = http.DefaultClient
	}

	if r.c.retryTimes == 0 {
		r.c.retryTimes = 1
	}

	var rawResp *http.Response
	// if meet error, retry times that you set
	for k := 0; k < r.c.retryTimes; k++ {
		rawResp, err = doRequest(r.c.client, request)
		if err != nil {
			// sleep retry delay
			time.Sleep(r.c.retryDelay)
			continue
		}
		break
	}

	if err != nil {
		return nil, err
	}

	if rawResp == nil {
		return nil, errors.New("empty resp")
	}

	if rawResp.StatusCode != 200 {
		return nil, errors.Errorf("unhealthy status code: [%d], status message: [%s]", rawResp.StatusCode, rawResp.Status)
	}

	return rawResp.Body, nil
}

func (r Result) TransformResponse() ([]byte, error) {
	if r.err != nil {
		return nil, r.err
	}

	// parse response data
	// code message data
	j, err := simplejson.NewJson(r.body)
	if err != nil {
		return nil, err
	}
	code, err := j.Get("code").Int()
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		message, _ := j.Get("message").String()
		return nil, fmt.Errorf(message)
	}
	marshalJSON, err := j.Get("data").MarshalJSON()
	if err != nil {
		return nil, err
	}
	return marshalJSON, nil
}

func (r Result) RawResponse() ([]byte, error) {
	return r.body, r.err
}

// Error returns the error executing the request, nil if no error occurred.
func (r Result) Error() error {
	return r.err
}
